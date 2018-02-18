// +build integration

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
	"testing"

	"github.com/chvck/meal-planner/model/planner"

	"github.com/shopspring/decimal"

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

// PLANNER TESTS

func TestPlannerAll(t *testing.T) {
	beforeEach(t)
	allPlanners := *testhelper.HelperCreatePlanners(t, sqlDb, "./testdata/planners.json")

	planners, err := planner.All(&adapter, 1517443200, 1519862400, 1)
	expectedPlanners := []planner.Planner{allPlanners[1], allPlanners[2], allPlanners[3]}
	assert.Nil(t, err)
	assert.Equal(t, 3, len(*planners))
	assert.Equal(t, expectedPlanners, *planners)
}

func TestPlannerAllOnlyWithinDateRange(t *testing.T) {
	beforeEach(t)
	allPlanners := *testhelper.HelperCreatePlanners(t, sqlDb, "./testdata/planners.json")

	planners, err := planner.All(&adapter, 1517443200, 1518998400, 1)
	expectedPlanners := []planner.Planner{allPlanners[1], allPlanners[2]}
	assert.Nil(t, err)
	assert.Equal(t, 2, len(*planners))
	assert.Equal(t, expectedPlanners, *planners)
}

func TestPlannerAddMenuNewPlanner(t *testing.T) {
	beforeEach(t)

	row := adapter.QueryOne(`INSERT INTO "menu" ("name", "description", "user_id") VALUES ($1, $2, $3) RETURNING id;`,
		"test", "test", 1)

	var menuID int
	if err := row.Scan(&menuID); err != nil {
		t.Fatal(err)
	}

	expectedWhen := 1517443200
	expectedFor := "breakfast"
	if err := planner.AddMenu(&adapter, expectedWhen, expectedFor, menuID, 1); err != nil {
		t.Fatal(err)
	}

	rows, err := sqlDb.Query(`SELECT id, "when", "for", "user_id" from planner;`)
	if err != nil {
		t.Fatal(err)
	}

	var planners []planner.Planner
	defer rows.Close()
	for rows.Next() {
		p := planner.Planner{}
		rows.Scan(&p.ID, &p.When, &p.For, &p.UserID)

		planners = append(planners, p)
	}

	assert.Equal(t, 1, len(planners))
	p := planners[0]

	assert.Equal(t, expectedWhen, p.When)
	assert.Equal(t, expectedFor, p.For)
	assert.Equal(t, 1, p.UserID)

	rows, err = sqlDb.Query(`SELECT "planner_id", "menu_id" from planner_to_menu;`)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var actualMenuId int
		var actualPlannerId int
		rows.Scan(&actualPlannerId, &actualMenuId)

		assert.Equal(t, menuID, actualMenuId)
		assert.Equal(t, p.ID, actualPlannerId)
	}
}

func TestPlannerAddMenuExistingPlanner(t *testing.T) {
	beforeEach(t)

	row := adapter.QueryOne(`INSERT INTO "menu" ("name", "description", "user_id") VALUES ($1, $2, $3) RETURNING id;`,
		"test", "test", 1)

	var menuID int
	if err := row.Scan(&menuID); err != nil {
		t.Fatal(err)
	}

	expectedWhen := 1517443200
	expectedFor := "breakfast"
	if _, err := sqlDb.Exec(`INSERT INTO "planner" ("when", "for", "user_id") VALUES ($1, $2, $3);`,
		expectedWhen, expectedFor, 1); err != nil {
		t.Fatal(err)
	}

	if err := planner.AddMenu(&adapter, expectedWhen, expectedFor, menuID, 1); err != nil {
		t.Fatal(err)
	}

	rows, err := sqlDb.Query(`SELECT id, "when", "for", "user_id" from planner;`)
	if err != nil {
		t.Fatal(err)
	}

	var planners []planner.Planner
	defer rows.Close()
	for rows.Next() {
		p := planner.Planner{}
		rows.Scan(&p.ID, &p.When, &p.For, &p.UserID)

		planners = append(planners, p)
	}

	assert.Equal(t, 1, len(planners))
	p := planners[0]

	assert.Equal(t, expectedWhen, p.When)
	assert.Equal(t, expectedFor, p.For)
	assert.Equal(t, 1, p.UserID)

	rows, err = sqlDb.Query(`SELECT "planner_id", "menu_id" from planner_to_menu;`)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var actualMenuId int
		var actualPlannerId int
		rows.Scan(&actualPlannerId, &actualMenuId)

		assert.Equal(t, menuID, actualMenuId)
		assert.Equal(t, p.ID, actualPlannerId)
	}
}

func TestPlannerAddMenuInvalidMealtime(t *testing.T) {
	beforeEach(t)
	row := adapter.QueryOne(`INSERT INTO "menu" ("name", "description", "user_id") VALUES ($1, $2, $3) RETURNING id;`,
		"test", "test", 1)

	var menuID int
	if err := row.Scan(&menuID); err != nil {
		t.Fatal(err)
	}

	expectedWhen := 1517443200
	expectedFor := "supper"

	err := planner.AddMenu(&adapter, expectedWhen, expectedFor, menuID, 1)

	assert.NotNil(t, err)
}

// MENU TESTS

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

// RECIPE TESTS

func TestRecipeOneWhenCorrectUserAndIdThenOK(t *testing.T) {
	beforeEach(t)

	recipes := *testhelper.HelperCreateRecipes(t, sqlDb, "./testdata/recipes.json")
	ingredients := *testhelper.HelperCreateIngredients(t, sqlDb, "./testdata/ingredients.json")

	expected := recipes[1]
	expected.Ingredients = ingredients[1]

	recipe, err := recipe.One(&adapter, expected.ID, 1)

	assert.Nil(t, err)
	assert.Equal(t, expected, *recipe)
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
		assert.Equal(t, expected, recipe)
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
	assert.Equal(t, expected, actual)
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
	assert.Equal(t, expected, actual)
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
		assert.Equal(t, expected, recipe)
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
		assert.Equal(t, expected, recipe)
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

func TestRecipeUpdate(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}
	r.ID = 1

	if _, err := sqlDb.Exec(
		`INSERT INTO "recipe" (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
		r.ID, r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.UserID); err != nil {
		t.Fatal(err)
	}

	for i, ing := range r.Ingredients {
		ing.ID = i
		if _, err := sqlDb.Exec(
			`INSERT INTO "ingredient" (id, name, measure, quantity, recipe_id) VALUES ($1, $2, $3, $4, $5);`,
			ing.ID, ing.Name, ing.Measure, ing.Quantity, r.ID,
		); err != nil {
			t.Fatal(err)
		}
	}

	newIng1 := ingredient.Ingredient{
		Name:     "Another ingredient",
		Quantity: decimal.NewFromFloat(4),
	}
	newIngredients := []ingredient.Ingredient{newIng1}

	r.Name = "Another recipe"
	r.Instructions = "Another set of instructions"
	r.Ingredients = newIngredients

	if err := recipe.Update(&adapter, r, r.UserID); err != nil {
		t.Fatal(err)
	}

	actualRecipe := recipe.Recipe{}
	row := sqlDb.QueryRow(`SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
		FROM recipe r where r.id = $1;`, r.ID)

	err = row.Scan(&actualRecipe.ID, &actualRecipe.Name, &actualRecipe.Instructions, &actualRecipe.Description, &actualRecipe.Yield,
		&actualRecipe.PrepTime, &actualRecipe.CookTime, &actualRecipe.UserID)
	if err != nil {
		t.Fatal(err)
	}

	var actualIngredients []ingredient.Ingredient
	rows, err := sqlDb.Query(fmt.Sprintf(`SELECT i.id, i.name, i.measure, i.quantity, i.recipe_id
		FROM ingredient i WHERE i.recipe_id = %v;`, r.ID))
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

	assert.Equal(t, 1, len(actualIngredients))
	assert.Equal(t, "Another ingredient", actualIngredients[0].Name)
	assert.Equal(t, null.String{NullString: sql.NullString{}}, actualIngredients[0].Measure)
	assert.Equal(t, decimal.NewFromFloat(4), actualIngredients[0].Quantity)
}

func TestRecipeUpdateWhenEmptyNameThenError(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}
	r.ID = 1

	if _, err := sqlDb.Exec(
		`INSERT INTO "recipe" (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
		r.ID, r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.UserID); err != nil {
		t.Fatal(err)
	}

	for i, ing := range r.Ingredients {
		ing.ID = i
		if _, err := sqlDb.Exec(
			`INSERT INTO "ingredient" (id, name, measure, quantity, recipe_id) VALUES ($1, $2, $3, $4, $5);`,
			ing.ID, ing.Name, ing.Measure, ing.Quantity, r.ID,
		); err != nil {
			t.Fatal(err)
		}
	}

	r.Name = ""

	err = recipe.Update(&adapter, r, 1)

	assert.NotNil(t, err)
}

func TestRecipeUpdateWhenEmptyInstructionsThenError(t *testing.T) {
	beforeEach(t)

	f := testhelper.HelperLoadFixture(t, "./testdata/create_recipe.json")
	var r recipe.Recipe
	err := json.Unmarshal(f, &r)
	if err != nil {
		t.Fatal(err)
	}
	r.ID = 1

	if _, err := sqlDb.Exec(
		`INSERT INTO "recipe" (id, name, instructions, yield, prep_time, cook_time, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`,
		r.ID, r.Name, r.Instructions, r.Yield, r.PrepTime, r.CookTime, r.Description, r.UserID); err != nil {
		t.Fatal(err)
	}

	for i, ing := range r.Ingredients {
		ing.ID = i
		if _, err := sqlDb.Exec(
			`INSERT INTO "ingredient" (id, name, measure, quantity, recipe_id) VALUES ($1, $2, $3, $4, $5);`,
			ing.ID, ing.Name, ing.Measure, ing.Quantity, r.ID,
		); err != nil {
			t.Fatal(err)
		}
	}

	r.Instructions = ""

	err = recipe.Update(&adapter, r, 1)

	assert.NotNil(t, err)
}

// INGREDIENT TESTS

func TestIngredientString(t *testing.T) {
	i := ingredient.Ingredient{
		ID:       1,
		Measure:  null.String{NullString: sql.NullString{String: "tbsp"}},
		Name:     "Paprika",
		Quantity: decimal.NewFromFloat(2),
		RecipeID: 2,
	}

	assert.Equal(t, "2 tbsp Paprika", fmt.Sprint(i))
}

// OTHER

func beforeEach(t *testing.T) {
	testhelper.HelperCleanDownModels(t, sqlDb)
	testhelper.HelperCreateUsers(t, sqlDb, "./testdata/users.json")
}
