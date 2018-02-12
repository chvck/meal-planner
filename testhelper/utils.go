package testhelper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/model/menu"
	"github.com/chvck/meal-planner/model/recipe"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"

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

// HelperCreateMenus writes menus + nested recipes + ingredients to the provided database using the fixtures at the path provided
func HelperCreateMenus(t *testing.T, db *sql.DB, path string) *map[int]menu.Menu {
	bytes := HelperLoadFixture(t, path)
	var menus []menu.Menu
	if err := json.Unmarshal(bytes, &menus); err != nil {
		t.Fatal(err)
	}

	idToMenu := make(map[int]menu.Menu)
	for _, m := range menus {
		query := `INSERT INTO "menu" (id, "name", "description", "user_id")
		VALUES ($1, $2, $3, $4)`
		if _, err := db.Exec(query, m.ID, m.Name, m.Description, m.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		idToMenu[m.ID] = m

		for _, rec := range m.Recipes {
			query := `INSERT INTO "recipe" (id, "name", "instructions", "yield", "prep_time", "cook_time", "description", "user_id")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			if _, err := db.Exec(query, rec.ID, rec.Name, rec.Instructions, rec.Yield, rec.PrepTime,
				rec.CookTime, rec.Description, rec.UserID); err != nil {
				t.Error(query)
				t.Fatal(err)
			}

			for _, ing := range rec.Ingredients {
				query := `INSERT INTO "ingredient" (id, "name", "quantity", "measure", "recipe_id")
				VALUES ($1, $2, $3, $4, $5)`
				if _, err := db.Exec(query, ing.ID, ing.Name, ing.Quantity, ing.Measure, ing.RecipeID); err != nil {
					t.Error(query)
					t.Fatal(err)
				}
			}
		}

	}

	return &idToMenu
}

// HelperCleanDownModels deletes from all model tables
func HelperCleanDownModels(t *testing.T, db *sql.DB) {
	if _, err := db.Exec(`DELETE FROM "ingredient"`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`DELETE FROM "recipe"`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`DELETE FROM "menu"`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`DELETE FROM "user"`); err != nil {
		t.Fatal(err)
	}
}

// HelperLoadFixture loads a fixture from a file
func HelperLoadFixture(t *testing.T, path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return bytes
}

/* The following two functions aren't performed within 1 connection for a couple of reasons.
A main driver is this issue https://github.com/mattes/migrate/issues/297 so this is an attempt
to address it as it's something being experienced in the tests.
*/

// HelperDatabaseConnection creates and returns a db connection
func HelperDatabaseConnection() (*sql.DB, string, func()) {
	cfg, err := config.Load("../config.test.json")
	if err != nil {
		panic(err)
	}

	openDb, err := sql.Open(cfg.DbType, cfg.DbString)
	if err != nil {
		panic(err)
	}

	teardown := func() {
		if err := openDb.Close(); err != nil {
			panic(err)
		}
	}

	return openDb, cfg.DbType, teardown
}

// HelperMigrate runs the database migrations
func HelperMigrate() {
	cfg, err := config.Load("../../config.test.json")
	if err != nil {
		panic(err)
	}

	m, err := migrate.New("file://../../migrations/", cfg.DbString)
	if err != nil {
		panic(err)
	}

	defer m.Close()

	if err := m.Down(); err != nil {
		fmt.Println(err)
	}

	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
}
