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

type menuToRecipes struct {
	MenuID    int   `json:"menu_id"`
	RecipeIDs []int `json:"recipe_ids"`
}

type plannerToMenus struct {
	PlannerID int   `json:"planner_id"`
	MenuIDs   []int `json:"menu_ids"`
}

type userWithPassword struct {
	model.User
	Password string `json:"password"`
}

type seed struct {
	Users            []userWithPassword `json:"users"`
	Planners         []model.Planner    `json:"planners"`
	Menus            []model.Menu       `json:"menus"`
	Recipes          []model.Recipe     `json:"recipes"`
	PlannerToRecipes plannerToRecipes   `json:"planner_to_recipes"`
	MenuToRecipes    menuToRecipes      `json:"menu_to_recipes"`
	PlannerToMenus   plannerToMenus     `json:"planner_to_menus"`
}

var sqlDb *sqlx.DB
var fixtures seed
var defaultUser model.User
var address string

func TestMain(m *testing.M) {
	cfg := loadConfig("../config.test.json")
	openDb, teardown := databaseConnection(cfg)
	defer teardown()
	sqlDb = openDb
	migrations(cfg)
	loadSeeds()
	defaultUser = fixtures.Users[0].User
	address = fmt.Sprintf("http://%v:%v/", cfg.Hostname, cfg.HTTPPort)
	srv, err := server.Run(cfg)
	if err != nil {
		panic(err)
	}
	defer srv.Shutdown(nil)

	code := m.Run()
	os.Exit(code)
}

func TestUserAuthorizationWhenUserThenOK(t *testing.T) {
	beforeEach(t)

	url := address + "testuser/"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserAuthorizationWhenNoAuthThenNotOK(t *testing.T) {
	beforeEach(t)

	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhenInvalidAuthThenNotOK(t *testing.T) {
	beforeEach(t)

	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "Bearer token", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhen1PartTokenThenNotOK(t *testing.T) {
	beforeEach(t)

	url := address + "testuser/"
	resp := sendRequest(t, "GET", url, "token", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserAuthorizationWhenAdminThenOK(t *testing.T) {
	beforeEach(t)

	url := address + "testuser/"
	token := createToken(&defaultUser, 2)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestAdminAuthorizationWhenAdminThenOK(t *testing.T) {
	beforeEach(t)

	url := address + "testadmin/"
	token := createToken(&defaultUser, 2)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestAdminAuthorizationWhenUserThenNotOK(t *testing.T) {
	beforeEach(t)

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
	bytes := loadFixture("testdata/seed.json")
	if err := json.Unmarshal(bytes, &fixtures); err != nil {
		panic(err)
	}
}

func beforeEach(t *testing.T) {
	cleanDownModels(t)
	createUsers(t, fixtures.Users)
	createPlanners(t, fixtures.Planners)
	createMenus(t, fixtures.Menus)
	createRecipes(t, fixtures.Recipes)
	createRelations(t, fixtures.PlannerToMenus, fixtures.PlannerToRecipes, fixtures.MenuToRecipes)
}

func createRelations(t *testing.T, pToM plannerToMenus, pToR plannerToRecipes, mToR menuToRecipes) {
	for _, mID := range pToM.MenuIDs {
		query := `INSERT INTO "planner_to_menu" ("planner_id", "menu_id")
		VALUES (?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, pToM.PlannerID, mID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}
	}

	for _, rID := range pToR.RecipeIDs {
		query := `INSERT INTO "planner_to_recipe" ("planner_id", "recipe_id")
		VALUES (?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, pToR.PlannerID, rID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}
	}

	for _, rID := range mToR.RecipeIDs {
		query := `INSERT INTO "menu_to_recipe" ("menu_id", "recipe_id")
		VALUES (?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, mToR.MenuID, rID); err != nil {
			t.Error(query)
			t.Fatal(err)
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

func createMenus(t *testing.T, menus []model.Menu) {
	for _, m := range menus {
		query := `INSERT INTO "menu" (id, "name", "description", "user_id")
		VALUES (?, ?, ?, ?)`
		query = sqlDb.Rebind(query)
		if _, err := sqlDb.Exec(query, m.ID, m.Name, m.Description, m.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
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
		if _, err := sqlDb.Exec(query, user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt, user.LastLogin, user.Password); err != nil {
			t.Error(query)
			t.Fatal(err)
		}
	}
}

func cleanDownModels(t *testing.T) {
	if _, err := sqlDb.Exec(`DELETE FROM "recipe"`); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlDb.Exec(`DELETE FROM "menu"`); err != nil {
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
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
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
		panic(err)
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
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return tokenString
}
