package server_test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"testing"

	"gopkg.in/guregu/null.v3"

	"github.com/chvck/meal-planner/controller"
	"github.com/chvck/meal-planner/model"
	"github.com/stretchr/testify/assert"
)

type MenuRecipe struct {
	MenuID   int `db:"menu_id"`
	RecipeID int `db:"recipe_id"`
}

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

func TestCreateMenu(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/"
	token := createToken(&defaultUser, 1)

	m := model.Menu{
		Name:        "test create",
		Description: null.StringFrom("Test description"),
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('menu_id_seq', (SELECT MAX(id) from menu));")
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

	var actual model.Menu
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	m.UserID = defaultUser.ID

	menusEqual(t, m, actual)

	dbMenu := menuFromDb(t, actual.ID)

	menusEqual(t, m, *dbMenu)
}

func TestCreateMenuWhenEmptyNameThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "menu/"
	token := createToken(&defaultUser, 1)

	r := model.Recipe{
		Name: "",
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlDb.Exec("SELECT setval('menu_id_seq', (SELECT MAX(id) from menu));")
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "POST", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var msgErr controller.JSONError
	json.Unmarshal(bodyBytes, &msgErr)

	assert.Equal(t, 400, resp.StatusCode)
	assert.NotNil(t, msgErr.Errors)
	assert.Len(t, msgErr.Errors, 1)
	assert.Equal(t, "name cannot be empty", msgErr.Errors[0])
}

func TestCreateMenuWhenNoAuthorizationThenError(t *testing.T) {
	cleanDownModels(t)

	url := address + "menu/"
	resp := sendRequest(t, "POST", url, "", nil)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUpdateMenu(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1"
	token := createToken(&defaultUser, 1)

	m := fixtures.Menus[1]
	m.Name = "updated name"
	m.Description = null.StringFrom("updated description")

	bytes, err := json.Marshal(m)
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

	var actual model.Menu
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	menusEqual(t, m, actual)

	dbMenu := menuFromDb(t, actual.ID)

	menusEqual(t, m, *dbMenu)
}

func TestUpdateMenuWhenNoNameThenError(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1"
	token := createToken(&defaultUser, 1)

	m := fixtures.Menus[1]
	m.Name = ""

	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	resp := sendRequest(t, "PUT", url, "Bearer "+token, bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var msgErr controller.JSONError
	json.Unmarshal(bodyBytes, &msgErr)

	assert.NotNil(t, msgErr.Errors)
	assert.Len(t, msgErr.Errors, 1)
	assert.Equal(t, "name cannot be empty", msgErr.Errors[0])
}

func TestAddRecipeMenu(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1/recipe/3"
	token := createToken(&defaultUser, 1)

	resp := sendRequest(t, "PUT", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	assert.Len(t, menuRecipesForMenu(t, 1), 3)
}

func TestRemoveRecipeMenu(t *testing.T) {
	opts := newResetOptions()
	opts.recreatePlanners = false
	resetDatabase(t, *opts)

	url := address + "menu/1/recipe/1"
	token := createToken(&defaultUser, 1)

	resp := sendRequest(t, "DELETE", url, "Bearer "+token, nil)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)

	assert.Len(t, menuRecipesForMenu(t, 1), 1)
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

func menuFromDb(t *testing.T, id int) *model.Menu {
	query := `SELECT m.id, m.name, m.description, m.user_id
	FROM menu m
	WHERE m.id = ?`

	query = sqlDb.Rebind(query)
	row := sqlDb.QueryRow(query, id)

	var dbMenu model.Menu
	if err := row.Scan(&dbMenu.ID, &dbMenu.Name, &dbMenu.Description, &dbMenu.UserID); err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		t.Fatal(err)
	}

	return &dbMenu
}

// Assert that two menus are equal if one doesn't have an ID
func menusEqual(t *testing.T, expected model.Menu, actual model.Menu) {
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.UserID, actual.UserID)
}

func menuRecipesForMenu(t *testing.T, menuID int) []MenuRecipe {
	query := `SELECT mr.menu_id, mr.recipe_id FROM menu_to_recipe mr WHERE mr.menu_id=?`
	query = sqlDb.Rebind(query)
	rows, err := sqlDb.Query(query, "1")
	if err != nil {
		t.Fatal(err)
	}

	menuRecipes := []MenuRecipe{}
	defer rows.Close()
	for rows.Next() {
		var menuRecipe MenuRecipe
		if err := rows.Scan(&menuRecipe.MenuID, &menuRecipe.RecipeID); err != nil {
			t.Fatal(err)
		}

		menuRecipes = append(menuRecipes, menuRecipe)
	}

	return menuRecipes
}
