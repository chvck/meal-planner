package testhelper

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/model/recipe"

	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/model/user"
)

type userWithPassword struct {
	user.User
	Password string `json:"password"`
}

// HelperCreateUsers writes users to the provided database using the fixtures at the path provided
func HelperCreateUsers(t *testing.T, db *sql.DB, path string) *map[int]user.User {
	bytes := helperLoadFixture(t, path)
	var users []userWithPassword
	if err := json.Unmarshal(bytes, &users); err != nil {
		t.Fatal(err)
	}

	userIDToUser := make(map[int]user.User)
	for _, user := range users {
		query := `INSERT INTO "user" (id, "username", "email", "created_at", "updated_at", "last_login", "password")
		VALUES (?, ?, ?, ?, ?, ?, ?)`
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
	bytes := helperLoadFixture(t, path)
	var ingredients []ingredient.Ingredient
	if err := json.Unmarshal(bytes, &ingredients); err != nil {
		t.Fatal(err)
	}

	idToIng := make(map[int][]ingredient.Ingredient)
	for _, ing := range ingredients {
		query := `INSERT INTO "ingredient" (id, "name", "quantity", "measure", "recipe_id")
		VALUES (?, ?, ?, ?, ?)`
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
	bytes := helperLoadFixture(t, path)
	var recipes []recipe.Recipe
	if err := json.Unmarshal(bytes, &recipes); err != nil {
		t.Fatal(err)
	}

	idToRecipe := make(map[int]recipe.Recipe)
	for _, rec := range recipes {
		query := `INSERT INTO "recipe" (id, "name", "instructions", "yield", "prep_time", "cook_time", "description", "user_id")
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
		if _, err := db.Exec(query, rec.ID, rec.Name, rec.Instructions, rec.Yield.Int64, rec.PrepTime,
			rec.CookTime, rec.Description, rec.UserID); err != nil {
			t.Error(query)
			t.Fatal(err)
		}

		idToRecipe[rec.ID] = rec
	}

	return &idToRecipe
}

func helperLoadFixture(t *testing.T, path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return bytes
}
