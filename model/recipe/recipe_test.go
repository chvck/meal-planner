package recipe_test

import (
	"database/sql"
	"encoding/json"
	"sort"
	"testing"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/testhelper"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/model/recipe"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*sql.DB, string, func()) {
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

	if err := m.Up(); err != nil {
		t.Fatal(err)
	}

	testhelper.HelperCreateUsers(t, openDb, "../testdata/users.json")

	down := func() {
		m.Down()
		openDb.Close()
	}

	return openDb, cfg.DbType, down
}

// -- Tests

func TestOneWhenCorrectUserAndIdThenOK(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	recipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, 1)

	assert.Nil(t, err)
	assertRecipe(t, &expected, recipe)
}

func TestOneWhenWrongUserThenNil(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
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

func TestOneWhenWrongIdThenNil(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	_ = *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipe, err := recipe.One(&adapter, 100, 1)

	assert.Nil(t, recipe)
	assert.Nil(t, err)
}

func TestAllWithLimit(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	expectedRecipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestAllWithLimitWhenNoResultsThenEmptySlice(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	_ = *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 0, 10)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(*recipes))
}

func TestAllWithLimitWhenLimitThenLimitedResults(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	expectedRecipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 1, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*recipes))
	actual := (*recipes)[0]
	expected := expectedRecipes[1]
	expected.Ingredients = ingredients[1]
	assertRecipe(t, &expected, &actual)
}

func TestAllWithLimitWhenOffsetThenOffsetResults(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	expectedRecipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*recipes))
	actual := (*recipes)[0]
	expected := expectedRecipes[2]
	expected.Ingredients = ingredients[2]
	assertRecipe(t, &expected, &actual)
}

func TestFindByIngredientNames1Name(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	expectedRecipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Paprika")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestFindByIngredientNamesWhenMultipleNamesThenOr(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	expectedRecipes := *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Chicken breast", "Potato")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestFindByIngredientNamesWhenNoResultsThenEmptySlice(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	_ = *testhelper.HelperCreateRecipes(t, openDb, "../testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, openDb, "../testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Fish")

	assert.Nil(t, err)
	assert.Equal(t, 0, len(*recipes))
}

func TestCreate(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	f := testhelper.HelperLoadFixture(t, "../testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}

	id, err := recipe.Create(&adapter, r, 1)
	if err != nil {
		t.Fatal(err)
	}

	actualRecipe := recipe.Recipe{}
	row := openDb.QueryRow(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r where r.id = $1;`, *id)

	err = row.Scan(&actualRecipe.ID, &actualRecipe.Name, &actualRecipe.Instructions, &actualRecipe.Description, &actualRecipe.Yield,
		&actualRecipe.PrepTime, &actualRecipe.CookTime, &actualRecipe.UserID)
	if err != nil {
		t.Fatal(err)
	}

	var actualIngredients []ingredient.Ingredient
	rows, err := openDb.Query(`SELECT i.id, i.name, i.measure, i.quantity, i.recipe_id
		FROM ingredient i WHERE i.recipe_id = 1;`)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		i := ingredient.Ingredient{}
		rows.Scan(&i.ID, &i.Name, &i.Measure, &i.Quantity, &i.RecipeID)

		actualIngredients = append(actualIngredients, i)
	}
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, r.Name, actualRecipe.Name)
	assert.Equal(t, r.CookTime, actualRecipe.CookTime)
	assert.Equal(t, r.PrepTime, actualRecipe.PrepTime)
	assert.Equal(t, r.Yield, actualRecipe.Yield)
	assert.Equal(t, r.Description, actualRecipe.Description)
	assert.Equal(t, r.Instructions, actualRecipe.Instructions)

	// Work out a way to ensure that the ingredients are actually created correctly
	assert.Equal(t, len(r.Ingredients), len(actualIngredients))
	assert.Equal(t, r.Ingredients[0].Name, actualIngredients[0].Name)
	assert.Equal(t, r.Ingredients[0].Measure, actualIngredients[0].Measure)
	assert.Equal(t, r.Ingredients[0].Quantity, actualIngredients[0].Quantity)
	assert.Equal(t, r.Ingredients[1].Name, actualIngredients[1].Name)
	assert.Equal(t, r.Ingredients[1].Measure, actualIngredients[1].Measure)
	assert.Equal(t, r.Ingredients[1].Quantity, actualIngredients[1].Quantity)
}

func TestCreateWhenEmptyNameThenError(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	f := testhelper.HelperLoadFixture(t, "../testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}
	r.Name = ""

	id, err := recipe.Create(&adapter, r, 1)

	assert.NotNil(t, err)
	assert.Nil(t, id)
}

func TestCreateWhenEmptyInstructionsThenError(t *testing.T) {
	openDb, dbType, teardown := setup(t)
	defer teardown()

	adapter := db.SqlxAdapter{}

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		t.Fatal(err)
		return
	}

	f := testhelper.HelperLoadFixture(t, "../testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}
	r.Instructions = ""

	id, err := recipe.Create(&adapter, r, 1)

	assert.NotNil(t, err)
	assert.Nil(t, id)
}

func assertRecipe(t *testing.T, expected *recipe.Recipe, actual *recipe.Recipe) {
	assert.NotNil(t, actual)
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
