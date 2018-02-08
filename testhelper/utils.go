package testhelper

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/model/recipe"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"

	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/model/user"
)

type userWithPassword struct {
	user.User
	Password string `json:"password"`
}

// HelperCreateUsers writes users to the provided database using the fixtures at the path provided
func HelperCreateUsers(t *testing.T, db *sql.DB, path string) *map[int]user.User {
	bytes := HelperLoadFixture(t, path)
	var users []userWithPassword
	if err := json.Unmarshal(bytes, &users); err != nil {
		t.Fatal(err)
	}

	userIDToUser := make(map[int]user.User)
	for _, user := range users {
		query := `INSERT INTO "user" (id, "username", "email", "created_at", "updated_at", "last_login", "password")
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
		if _, err := db.Exec(query, user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt, user.LastLogin.Int64, user.Password); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		userIDToUser[user.ID] = user.User
	}

	return &userIDToUser
}

// HelperCreateIngredients writes ingredients to the provided database using the fixtures at the path provided
func HelperCreateIngredients(t *testing.T, db *sql.DB, path string) *map[int][]ingredient.Ingredient {
	bytes := HelperLoadFixture(t, path)
	var ingredients []ingredient.Ingredient
	if err := json.Unmarshal(bytes, &ingredients); err != nil {
		t.Fatal(err)
	}

	idToIng := make(map[int][]ingredient.Ingredient)
	for _, ing := range ingredients {
		query := `INSERT INTO "ingredient" (id, "name", "quantity", "measure", "recipe_id")
		VALUES ($1, $2, $3, $4, $5)`
		if _, err := db.Exec(query, ing.ID, ing.Name, ing.Quantity, ing.Measure, ing.RecipeID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		idToIng[ing.RecipeID] = append(idToIng[ing.RecipeID], ing)
	}

	return &idToIng
}

// HelperCreateRecipes writes recipes to the provided database using the fixtures at the path provided
func HelperCreateRecipes(t *testing.T, db *sql.DB, path string) *map[int]recipe.Recipe {
	bytes := HelperLoadFixture(t, path)
	var recipes []recipe.Recipe
	if err := json.Unmarshal(bytes, &recipes); err != nil {
		t.Fatal(err)
	}

	idToRecipe := make(map[int]recipe.Recipe)
	for _, rec := range recipes {
		query := `INSERT INTO "recipe" (id, "name", "instructions", "yield", "prep_time", "cook_time", "description", "user_id")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		if _, err := db.Exec(query, rec.ID, rec.Name, rec.Instructions, rec.Yield, rec.PrepTime,
			rec.CookTime, rec.Description, rec.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		idToRecipe[rec.ID] = rec
	}

	return &idToRecipe
}

// HelperLoadFixture loads a fixture from a file
func HelperLoadFixture(t *testing.T, path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return bytes
}

// HelperSetupModels runs migrations, creates users and returns a db connection
func HelperSetupModels(t *testing.T) (*sql.DB, string, func()) {
	cfg, err := config.Load("../../config.test.json")
	if err != nil {
		t.Fatal(err)
	}

	openDb, err := sql.Open(cfg.DbType, cfg.DbString)
	if err != nil {
		t.Fatal(err)
	}

	driver, err := postgres.WithInstance(openDb, &postgres.Config{})
	if err != nil {
		t.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../migrations/", "postgres", driver)

	if err != nil {
		t.Fatal(err)
	}

	if err := m.Down(); err != nil {
		t.Fatal(err)
	}

	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	HelperCreateUsers(t, openDb, "../testdata/users.json")

	if err := openDb.Close(); err != nil {
		t.Fatal(err)
	}

	openDb, err = sql.Open(cfg.DbType, cfg.DbString)
	if err != nil {
		t.Fatal(err)
	}

	down := func() {
		openDb.Close()
	}

	return openDb, cfg.DbType, down
}
