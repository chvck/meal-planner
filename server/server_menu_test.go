package server_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAllMenus(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var menus []model.Menu
	err = json.Unmarshal(bodyBytes, &menus)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Menus[0:3]

	assert.Equal(t, expected, menus)
}

func TestGetAllMenusWhen2PerPageThen2(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/?perPage=2"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var menus []model.Menu
	err = json.Unmarshal(bodyBytes, &menus)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Menus[0:2]

	assert.Equal(t, expected, menus)
}

func TestGetAllMenusWhenOffsetThenOffset(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/?offset=2"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var menus []model.Menu
	err = json.Unmarshal(bodyBytes, &menus)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Menus[2:3]

	assert.Equal(t, expected, menus)
}

func TestGetAllMenusWhenNoneThenEmptyArray(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/?offset=5"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var menus []model.Menu
	err = json.Unmarshal(bodyBytes, &menus)
	if err != nil {
		t.Fatal(err)
	}

	expected := []model.Menu{}

	assert.Equal(t, expected, menus)
}

func TestGetAllMenusWithRecipes(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/?includeRecipes=true"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	menus := make([]model.Menu, 0)
	err = json.Unmarshal(bodyBytes, &menus)
	if err != nil {
		t.Fatal(err)
	}

	expectedMenus := make([]model.Menu, 3)
	copy(expectedMenus, fixtures.Menus[0:3])
	for i, m := range expectedMenus {
		recipes := recipesFromDbForMenuID(t, m.ID)
		expectedMenus[i].Recipes = recipes
	}

	assert.Equal(t, expectedMenus, menus)
}

func TestGetAllMenusWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "menu/"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestGetOneMenu(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual model.Menu
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Menus[0]

	assert.Equal(t, expected, actual)
}

func TestGetOneMenuWithRecipe(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1?includeRecipes=true"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual model.Menu
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	expected := fixtures.Menus[0]
	recipes := recipesFromDbForMenuID(t, expected.ID)
	expected.Recipes = recipes

	assert.Equal(t, expected, actual)
}

func TestGetOneMenuWhenBelongsToOtherUserThenNull(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/4"
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

func TestGetOneMenuWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "menu/1"
	resp := sendRequest(t, "GET", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestGetOneMenuWhenInvalidIDClientError(t *testing.T) {
	cleanDownModels(t)

	url := address + "menu/somethingwrong"
	token := createToken(&defaultUser, 1)
	resp := sendRequest(t, "GET", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func recipesFromDbForMenuID(t *testing.T, id int) []model.Recipe {
	query := `SELECT r.id, r.name, r.instructions, r.description, r.yield, r.prep_time, r.cook_time, r.user_id
	FROM recipe r
	JOIN menu_to_recipe mr ON mr.recipe_id = r.id
	WHERE mr.menu_id = ?
	ORDER by r.id`

	query = sqlDb.Rebind(query)
	rows, err := sqlDb.Query(query, id)
	if err != nil {
		t.Fatal(err)
	}
	recipes := []model.Recipe{}
	defer rows.Close()
	for rows.Next() {
		var dbRecipe model.Recipe
		if err := rows.Scan(&dbRecipe.ID, &dbRecipe.Name, &dbRecipe.Instructions, &dbRecipe.Description, &dbRecipe.Yield,
			&dbRecipe.PrepTime, &dbRecipe.CookTime, &dbRecipe.UserID); err != nil {
			t.Fatal(err)
		}

		ingredients := ingredientsFromDb(t, dbRecipe.ID)
		dbRecipe.Ingredients = ingredients

		recipes = append(recipes, dbRecipe)
	}

	return recipes
}
