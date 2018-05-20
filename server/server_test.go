package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/server"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

type plannerToRecipes struct {
	PlannerID int   `json:"planner_id"`
	RecipeIDs []int `json:"recipe_ids"`
}

type userWithPassword struct {
	model.User
	Password string `json:"password"`
}

type seed struct {
	Users            []userWithPassword `json:"users"`
	Planners         []model.Planner    `json:"planners"`
	Recipes          []model.Recipe     `json:"recipes"`
	PlannerToRecipes []plannerToRecipes `json:"planner_to_recipe"`
}

type resetOptions struct {
	recreateUsers    bool
	recreateRecipes  bool
	recreatePlanners bool
}

var sqlDb *sqlx.DB
var fixtures seed
var defaultUser model.User
var address string
var authKey string

func TestMain(m *testing.M) {
	cfg := loadConfig("../config.test.json")
	openDb, teardown := databaseConnection(cfg)
	defer teardown()
	sqlDb = openDb
	migrations(cfg)
	loadSeeds()
	defaultUser = fixtures.Users[0].User
	address = fmt.Sprintf("http://%v:%v/", cfg.Hostname, cfg.HTTPPort)
	authKey = cfg.AuthKey
	srv, err := server.Run(cfg)
	if err != nil {
		panic(err)
	}
	defer srv.Shutdown(nil)

	code := m.Run()
	os.Exit(code)
}

func TestUserAuthorizationWhenUserThenOK(t *testing.T) {
	url := address + "testuser/"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserAuthorizationWhenNoAuthThenNotOK(t *testing.T) {
	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhenInvalidAuthThenNotOK(t *testing.T) {
	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "Bearer token", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhen1PartTokenThenNotOK(t *testing.T) {
	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "token", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhenAdminThenOK(t *testing.T) {
	url := address + "testuser/"
	token := createToken(&defaultUser, 2)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestAdminAuthorizationWhenAdminThenOK(t *testing.T) {
	url := address + "testadmin/"
	token := createToken(&defaultUser, 2)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestAdminAuthorizationWhenUserThenNotOK(t *testing.T) {
	url := address + "testadmin/"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func sendRequest(t *testing.T, method string, url string, tok string, data []byte) *http.Response {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	if tok != "" {
		req.Header.Set("authorization", tok)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func loadSeeds() {
	b := loadFixture("testdata/seed.json")
	if err := json.Unmarshal(b, &fixtures); err != nil {
		panic(err)
	}
}

func newResetOptions() *resetOptions {
	return &resetOptions{
		recreatePlanners: true,
		recreateRecipes:  true,
		recreateUsers:    true,
	}
}

func resetDatabase(t *testing.T, opts resetOptions) {
	cleanDownModels(t)
	if opts.recreateUsers {
		createUsers(t, fixtures.Users)
	}
	if opts.recreatePlanners {
		createPlanners(t, fixtures.Planners)
	}
	if opts.recreateRecipes {
		createRecipes(t, fixtures.Recipes)
	}
	createRelations(t, fixtures.PlannerToRecipes, opts)
}

func createRelations(t *testing.T, pToR []plannerToRecipes, opts resetOptions) {
	if opts.recreatePlanners && opts.recreateRecipes {
		for _, rel := range pToR {
			for _, rID := range rel.RecipeIDs {
				query := `INSERT INTO "planner_to_recipe" ("planner_id", "recipe_id")
		VALUES (?, ?)`
				query = sqlDb.Rebind(query)
				if _, err := sqlDb.Exec(query, rel.PlannerID, rID); err != nil {
					t.Error(query)
					t.Fatal(err)
				}
			}
		}
	}
}

func createRecipes(t *testing.T, recipes []model.Recipe) {
	for _, rec := range recipes {
		query := `INSERT INTO "recipe" (id, "name", "instructions", "yield", "prep_time", "cook_time", "description", "user_id")
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, rec.ID, rec.Name, rec.Instructions, rec.Yield, rec.PrepTime,
			rec.CookTime, rec.Description, rec.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		for _, ing := range rec.Ingredients {
			query := `INSERT INTO "ingredient" (id, "name", "quantity", "measure", "recipe_id")
			VALUES (?, ?, ?, ?, ?)`
			query = sqlDb.Rebind(query)
			if _, err := sqlDb.Exec(query, ing.ID, ing.Name, ing.Quantity, ing.Measure, ing.RecipeID); err != nil {
				t.Error(query)
				t.Fatal(err)
			}
		}
	}
}

func createPlanners(t *testing.T, planners []model.Planner) {
	for _, p := range planners {
		query := `INSERT INTO "planner" (id, "when", "for", "user_id")
		VALUES (?, ?, ?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, p.ID, p.When, p.For, p.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}
	}
}

func createUsers(t *testing.T, users []userWithPassword) {
	for _, user := range users {
		query := `INSERT INTO "user" (id, "username", "email", "created_at", "updated_at", "last_login", "password")
		VALUES (?, ?, ?, ?, ?, ?, ?)`
		query = sqlDb.Rebind(query)

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		if _, err := sqlDb.Exec(query, user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt,
			user.LastLogin, hash); err != nil {
			t.Error(query)
			t.Fatal(err)
		}
	}
}

func cleanDownModels(t *testing.T) {
	if _, err := sqlDb.Exec(`DELETE FROM "recipe"`); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlDb.Exec(`DELETE FROM "planner"`); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlDb.Exec(`DELETE FROM "user"`); err != nil {
		t.Fatal(err)
	}
}

func loadFixture(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return b
}

func loadConfig(path string) *config.Info {
	cfg, err := config.Load(path)
	if err != nil {
		panic(err)
	}

	return cfg
}

func databaseConnection(cfg *config.Info) (*sqlx.DB, func()) {
	openDb, err := sqlx.Open(cfg.DbType, cfg.DbString)
	if err != nil {
		panic(err)
	}

	teardown := func() {
		if err := openDb.Close(); err != nil {
			panic(err)
		}
	}

	return openDb, teardown
}

// We let migrations handle its own db connection due to oddities with
// locks
func migrations(cfg *config.Info) {
	m, err := migrate.New("file://../migrations/", cfg.DbString)
	if err != nil {
		panic(err)
	}

	defer m.Close()

	if err := m.Down(); err != nil {
		fmt.Println("Couldn't perform migrations down")
	}

	if err := m.Up(); err != nil {
		panic(err)
	}
}

func createToken(user *model.User, level int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.ID,
		"lastLogin": user.LastLogin,
		"level":     level,
	})
	tokenString, err := token.SignedString([]byte(authKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}
