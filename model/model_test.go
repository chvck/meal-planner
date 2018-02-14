package model_test

/*
All model tests in one file due to issues related to acquiring locks with the ps driver.
See: https://github.com/mattes/migrate/issues/297
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/model/ingredient"
	"github.com/chvck/meal-planner/model/menu"
	"github.com/chvck/meal-planner/model/recipe"
	"github.com/chvck/meal-planner/testhelper"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"
)

var adapter db.SqlxAdapter
var sqlDb *sql.DB

func TestMain(m *testing.M) {
	openDb, dbType, teardown := testhelper.HelperDatabaseConnection()
	defer teardown()
	sqlDb = openDb
	testhelper.HelperMigrate()

	if err := adapter.InitializeWithDb(sqlx.NewDb(openDb, dbType)); err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestMenuOneWhenCorrectUserAndIdThenOK(t *testing.T) {
	beforeEach(t)
	menus := *testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	expected := menus[1]

	m, err := menu.One(&adapter, expected.ID, 1)

	assert.Nil(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, &expected, m)
}

func TestMenuOneWhenWrongUserAndIdThenNil(t *testing.T) {
	beforeEach(t)
	testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	m, err := menu.One(&adapter, 1, 2)

	assert.Nil(t, err)
	assert.Nil(t, m)
}

func TestMenuOneWhenWrongIdThenNil(t *testing.T) {
	beforeEach(t)
	testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	m, err := menu.One(&adapter, 100, 1)

	assert.Nil(t, err)
	assert.Nil(t, m)
}

func TestMenuAllWithLimit(t *testing.T) {
	beforeEach(t)

	allMenus := *testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	menus, err := menu.AllWithLimit(&adapter, 10, 0, 1)
	expectedMenus := []menu.Menu{allMenus[1], allMenus[2]}
	assert.Nil(t, err)
	assert.Equal(t, 2, len(*menus))
	assert.Equal(t, expectedMenus, *menus)
}

func TestMenuAllWithLimitWhenNoResultsThenEmptySlice(t *testing.T) {
	beforeEach(t)

	testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	menus, err := menu.AllWithLimit(&adapter, 10, 0, 10)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(*menus))
}

func TestMenuAllWithLimitWhenLimitThenLimitedResults(t *testing.T) {
	beforeEach(t)

	allMenus := *testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	menus, err := menu.AllWithLimit(&adapter, 1, 0, 1)
	expectedMenus := []menu.Menu{allMenus[1]}
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*menus))
	assert.Equal(t, expectedMenus, *menus)
}

func TestMenuAllWithLimitWhenOffsetThenOffsetResults(t *testing.T) {
	beforeEach(t)

	allMenus := *testhelper.HelperCreateMenus(t, sqlDb, "./testdata/menus.json")

	menus, err := menu.AllWithLimit(&adapter, 10, 1, 1)
	expectedMenus := []menu.Menu{allMenus[2]}
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*menus))
	assert.Equal(t, expectedMenus, *menus)
}

func TestRecipeOneWhenCorrectUserAndIdThenOK(t *testing.T) {
	beforeEach(t)

	recipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, 1)

	assert.Nil(t, err)
	assertRecipe(t, &expected, recipe)
}

func TestRecipeOneWhenWrongUserThenNil(t *testing.T) {
	beforeEach(t)

	recipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, 2)

	assert.Nil(t, recipe)
	assert.Nil(t, err)
}

func TestRecipeOneWhenWrongIdThenNil(t *testing.T) {
	beforeEach(t)

	_ = *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipe, err := recipe.One(&adapter, 100, 1)

	assert.Nil(t, recipe)
	assert.Nil(t, err)
}

func TestRecipeAllWithLimit(t *testing.T) {
	beforeEach(t)

	expectedRecipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestRecipeAllWithLimitWhenNoResultsThenEmptySlice(t *testing.T) {
	beforeEach(t)

	_ = *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 0, 10)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(*recipes))
}

func TestRecipeAllWithLimitWhenLimitThenLimitedResults(t *testing.T) {
	beforeEach(t)

	expectedRecipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 1, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*recipes))
	actual := (*recipes)[0]
	expected := expectedRecipes[1]
	expected.Ingredients = ingredients[1]
	assertRecipe(t, &expected, &actual)
}

func TestRecipeAllWithLimitWhenOffsetThenOffsetResults(t *testing.T) {
	beforeEach(t)

	expectedRecipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.AllWithLimit(&adapter, 10, 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(*recipes))
	actual := (*recipes)[0]
	expected := expectedRecipes[2]
	expected.Ingredients = ingredients[2]
	assertRecipe(t, &expected, &actual)
}

func TestRecipeFindByIngredientNames1Name(t *testing.T) {
	beforeEach(t)

	expectedRecipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Paprika")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestRecipeFindByIngredientNamesWhenMultipleNamesThenOr(t *testing.T) {
	beforeEach(t)

	expectedRecipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Chicken breast", "Potato")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(*recipes))
	for _, recipe := range *recipes {
		expected := expectedRecipes[recipe.ID]
		expected.Ingredients = ingredients[expected.ID]
		assertRecipe(t, &expected, &recipe)
	}
}

func TestRecipeFindByIngredientNamesWhenNoResultsThenEmptySlice(t *testing.T) {
	beforeEach(t)

	_ = *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	_ = *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	recipes, err := recipe.FindByIngredientNames(&adapter, "Fish")

	assert.Nil(t, err)
	assert.Equal(t, 0, len(*recipes))
}

func TestRecipeCreate(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
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
	row := sqlDb.QueryRow(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r where r.id = $1;`, *id)

	err = row.Scan(&actualRecipe.ID, &actualRecipe.Name, &actualRecipe.Instructions, &actualRecipe.Description, &actualRecipe.Yield,
		&actualRecipe.PrepTime, &actualRecipe.CookTime, &actualRecipe.UserID)
	if err != nil {
		t.Fatal(err)
	}

	var actualIngredients []ingredient.Ingredient
	rows, err := sqlDb.Query(fmt.Sprintf(`SELECT i.id, i.name, i.measure, i.quantity, i.recipe_id
		FROM ingredient i WHERE i.recipe_id = %v;`, *id))
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

func TestRecipeCreateWhenEmptyNameThenError(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
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

func TestRecipeCreateWhenEmptyInstructionsThenError(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
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

func TestIngredientString(t *testing.T) {
	i := ingredient.Ingredient{
		ID:       1,
		Measure:  null.String{NullString: sql.NullString{String: "tbsp"}},
		Name:     "Paprika",
		Quantity: 2,
		RecipeID: 2,
	}

	assert.Equal(t, "2 tbsp Paprika", fmt.Sprint(i))
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

func beforeEach(t *testing.T) {
	testhelper.HelperCleanDownModels(t, sqlDb)
	testhelper.HelperCreateUsers(t, sqlDb, "./testdata/users.json")
}
