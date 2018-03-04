package server_test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"
)

func TestGetAllRecipes(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var recipes []model.Recipe
	err = json.Unmarshal(bodyBytes, &recipes)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Recipes[0:4]

	assert.Equal(t, expected, recipes)
}

func TestGetAllRecipesWhen2PerPageThen2(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/?perPage=2"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var recipes []model.Recipe
	err = json.Unmarshal(bodyBytes, &recipes)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Recipes[0:2]

	assert.Equal(t, expected, recipes)
}

func TestGetAllRecipesWhenOffsetThenOffset(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/?offset=3"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var recipes []model.Recipe
	err = json.Unmarshal(bodyBytes, &recipes)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Recipes[3:4]

	assert.Equal(t, expected, recipes)
}

func TestGetAllRecipesWhenNoneThenEmptyArray(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/?offset=5"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var recipes []model.Recipe
	err = json.Unmarshal(bodyBytes, &recipes)
	if err != nil {
		t.Fatal(err)
	}

	expected := []model.Recipe{}

	assert.Equal(t, expected, recipes)
}

func TestGetAllRecipesWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestGetOneRecipe(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/1"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual model.Recipe
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Recipes[0]

	assert.Equal(t, expected, actual)
}

func TestGetOneRecipeWhenBelongsToOtherUserThenNull(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/5"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "null", string(bodyBytes))
}

func TestGetOneRecipeWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/1"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestGetOneRecipeWhenInvalidIDClientError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/somethingwrong"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateRecipe(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/"
	token := createToken(&defaultUser, 1)

	r := model.Recipe{
		Name:         "test create",
		Instructions: "test instructions",
		UserID:       1,
		Ingredients: []model.Ingredient{
			model.Ingredient{
				Name:     "ing1",
				Quantity: decimal.NewFromFloat(1),
				Measure:  null.StringFrom("meas 1"),
			},
			model.Ingredient{
				Name:     "ing2",
				Quantity: decimal.NewFromFloat(2),
			},
		},
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('recipe_id_seq', (SELECT MAX(id) from recipe));")
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('ingredient_id_seq', (SELECT MAX(id) from ingredient));")
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "POST", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 201, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual model.Recipe
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	recipesEqual(t, r, actual)
	ingredientsEqualAndCorrect(t, r.Ingredients, actual.Ingredients, actual.ID)

	dbRecipe := recipeFromDb(t, actual.ID)
	dbIngredients := ingredientsFromDb(t, actual.ID)

	recipesEqual(t, r, *dbRecipe)
	ingredientsEqualAndCorrect(t, r.Ingredients, dbIngredients, dbRecipe.ID)
}

func TestCreateRecipeWhenEmptyNameThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/"
	token := createToken(&defaultUser, 1)

	r := model.Recipe{
		Name:         "",
		Instructions: "test instructions",
		UserID:       1,
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('recipe_id_seq', (SELECT MAX(id) from recipe));")
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "POST", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateRecipeWhenEmptyInstructionsThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/"
	token := createToken(&defaultUser, 1)

	r := model.Recipe{
		Name:         "test name",
		Instructions: "",
		UserID:       1,
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('recipe_id_seq', (SELECT MAX(id) from recipe));")
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "POST", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestCreateRecipeWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/"
	resp := sendRequest(t, "POST", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUpdateRecipe(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/1"
	token := createToken(&defaultUser, 1)

	r := fixtures.Recipes[0]

	r.Name = "Updated name"
	r.Instructions = "Updated instructions"
	r.Description = null.StringFrom("Updated desc")
	r.Yield = null.IntFrom(20)
	r.Ingredients[0].Name = "Updated ing name"

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "PUT", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual model.Recipe
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	recipesEqual(t, r, actual)
	ingredientsEqualAndCorrect(t, r.Ingredients, actual.Ingredients, actual.ID)

	dbRecipe := recipeFromDb(t, actual.ID)
	dbIngredients := ingredientsFromDb(t, actual.ID)

	recipesEqual(t, r, *dbRecipe)
	ingredientsEqualAndCorrect(t, r.Ingredients, dbIngredients, dbRecipe.ID)
}

func TestUpdateRecipeWhenEmptyNameThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/1"
	token := createToken(&defaultUser, 1)

	r := fixtures.Recipes[0]

	r.Name = ""

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "PUT", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdateRecipeWhenEmptyInstructionsThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/1"
	token := createToken(&defaultUser, 1)

	r := fixtures.Recipes[0]

	r.Instructions = ""

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "PUT", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdateRecipeWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/1"
	resp := sendRequest(t, "PUT", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestDeleteRecipe(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "recipe/1"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "DELETE", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 204, resp.StatusCode)

	dbRecipe := recipeFromDb(t, 1)
	dbIngredients := ingredientsFromDb(t, 1)

	assert.Nil(t, dbRecipe)
	assert.Len(t, dbIngredients, 0)
}

func TestDeleteRecipeWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "recipe/1"
	resp := sendRequest(t, "DELETE", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

// Assert that two recipes are equal if one doesn't have an ID
func recipesEqual(t *testing.T, expected model.Recipe, actual model.Recipe) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.CookTime, actual.CookTime)
	assert.Equal(t, expected.PrepTime, actual.PrepTime)
	assert.Equal(t, expected.Instructions, actual.Instructions)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Yield, actual.Yield)
}

// Assert that two lists of ingredients are equal if one list doesn't have IDs,
// also checks that the actual ingredients have correct recipe IDs
func ingredientsEqualAndCorrect(t *testing.T, expected []model.Ingredient, actual []model.Ingredient, recipeID int) {
	assert.Equal(t, len(expected), len(actual))

	for i, r := range actual {
		expectedR := expected[i]
		assert.Equal(t, expectedR.Name, r.Name)
		assert.Equal(t, expectedR.Quantity, r.Quantity)
		assert.Equal(t, expectedR.Measure, r.Measure)
		assert.Equal(t, recipeID, r.RecipeID)
	}
}

func recipeFromDb(t *testing.T, id int) *model.Recipe {
	query := `SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
	FROM recipe r
	WHERE r.id = ?`

	query = sqlDb.Rebind(query)
	row := sqlDb.QueryRow(query, id)

	var dbRecipe model.Recipe
	if err := row.Scan(&dbRecipe.ID, &dbRecipe.Name, &dbRecipe.Instructions, &dbRecipe.Description, &dbRecipe.Yield,
		&dbRecipe.PrepTime, &dbRecipe.CookTime, &dbRecipe.UserID); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		t.Fatal(err)
	}

	return &dbRecipe
}

func ingredientsFromDb(t *testing.T, recipeID int) []model.Ingredient {
	query := `SELECT id, recipe_id, name, measure, quantity
		FROM ingredient
		WHERE recipe_id = ?
		ORDER BY id;`

	query = sqlDb.Rebind(query)

	ingredients := []model.Ingredient{}
	rows, err := sqlDb.Query(query, recipeID)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			rID     int
			ingID   int
			ingName string
			mName   null.String
			q       decimal.Decimal
		)
		if err := rows.Scan(&ingID, &rID, &ingName, &mName, &q); err != nil {
			t.Fatal(err)
		}

		i := model.Ingredient{ID: ingID, RecipeID: rID, Name: ingName, Measure: mName, Quantity: q}
		ingredients = append(ingredients, i)
	}

	if err = rows.Err(); err != nil {
		t.Fatal(err)
	}

	return ingredients
}
