package recipe_test

import (
	"database/sql"
	"sort"
	"testing"

	"github.com/chvck/meal-planner/testhelper"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/recipe"
	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*sql.DB, func()) {
	openDb, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatal(err)
	}

	driver, err := sqlite3.WithInstance(openDb, &sqlite3.Config{})

	if err != nil {
		t.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../migrations/", "sqlite3", driver)

	if err != nil {
		t.Fatal(err)
	}

	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	testhelper.HelperCreateUsers(t, openDb, "../testdata/users.json")

	down := func() {
		openDb.Close()
		m.Down()
	}

	return openDb, down
}

// -- Tests

func TestOneWhenCorrectUserAndIdThenOK(t *testing.T) {
	openDb, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, "sqlite3")); err != nil {
		t.Fatal(err)
		return
	}

	recipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, expected.UserID)

	assert.Nil(t, err)
	assertRecipe(t, &expected, recipe)
}

func TestOneWhenWrongUserThenNoResult(t *testing.T) {
	openDb, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, "sqlite3")); err != nil {
		t.Fatal(err)
		return
	}

	recipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, 2)

	assert.Nil(t, recipe)
	assert.Nil(t, err)
}

func assertRecipe(t *testing.T, expected *recipe.Recipe, actual *recipe.Recipe) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.CookTime, actual.CookTime)
	assert.Equal(t, expected.PrepTime, actual.PrepTime)
	assert.Equal(t, expected.Yield, actual.Yield)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Instructions, actual.Instructions)

	// Slices of different orders aren't equal
	sort.SliceStable(expected.Ingredients, func(i, j int) bool {
		return expected.Ingredients[i].ID < expected.Ingredients[j].ID
	})

	sort.SliceStable(actual.Ingredients, func(i, j int) bool {
		return actual.Ingredients[i].ID < actual.Ingredients[j].ID
	})
	assert.Equal(t, expected.Ingredients, actual.Ingredients)
}
