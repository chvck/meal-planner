package server_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/chvck/meal-planner/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestUserLogin(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	opts.recreateRecipes = false
	resetDatabase(t, *opts)

	user := fixtures.Users[0]

	creds := loginCredentials{Username: user.Username, Password: user.Password}
	bytes, err := json.Marshal(creds)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "login/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserLoginWhenUserDoesntExistThenError(t *testing.T) {
	opts := newResetOptions()
	opts.recreateMenus = false
	opts.recreatePlanners = false
	opts.recreateRecipes = false
	resetDatabase(t, *opts)

	creds := loginCredentials{Username: "wrong user", Password: "wrong password"}
	bytes, err := json.Marshal(creds)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "login/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 401, resp.StatusCode)
}

func TestUserCreate(t *testing.T) {
	cleanDownModels(t)

	u := model.User{
		Username: "user create",
		Email:    "test@test.test",
	}

	uWP := userWithPassword{
		u,
		"testpassword",
	}

	_, err := sqlDb.Exec(`SELECT setval('user_id_seq', (SELECT MAX(id) from "user"));`)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := json.Marshal(uWP)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "register/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 201, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var actual userWithPassword
	err = json.Unmarshal(bodyBytes, &actual)
	if err != nil {
		t.Fatal(err)
	}

	usersEqual(t, u, actual.User)
	assert.Zero(t, actual.Password)

	dbUser := userFromDb(t, actual.ID)

	usersEqual(t, u, dbUser.User)

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(uWP.Password))

	assert.Nil(t, err)
}

func TestUserCreateWhenNoUsernameThenError(t *testing.T) {
	cleanDownModels(t)

	u := model.User{
		Username: "",
		Email:    "test@test.test",
	}

	uWP := userWithPassword{
		u,
		"testpassword",
	}

	bytes, err := json.Marshal(uWP)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "register/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUserCreateWhenNoEmailThenError(t *testing.T) {
	cleanDownModels(t)

	u := model.User{
		Username: "test name",
		Email:    "",
	}

	uWP := userWithPassword{
		u,
		"testpassword",
	}

	bytes, err := json.Marshal(uWP)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "register/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUserCreateWhenPasswordLessThan8CharsThenError(t *testing.T) {
	cleanDownModels(t)

	u := model.User{
		Username: "test name",
		Email:    "test@test.test",
	}

	uWP := userWithPassword{
		u,
		`t\*tp$A`,
	}

	bytes, err := json.Marshal(uWP)
	if err != nil {
		t.Fatal(err)
	}

	url := address + "register/"
	resp := sendRequest(t, "POST", url, "", bytes)
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
}

// Assert that two users are equal if one doesn't have an ID
func usersEqual(t *testing.T, expected model.User, actual model.User) {
	assert.Equal(t, expected.Username, actual.Username)
	assert.Equal(t, expected.Email, actual.Email)
	assert.NotZero(t, actual.ID)
	assert.NotZero(t, actual.CreatedAt)
	assert.NotZero(t, actual.UpdatedAt)
	assert.Zero(t, actual.LastLogin)
}

func userFromDb(t *testing.T, id int) *userWithPassword {
	query := `SELECT id, username, email, password, created_at, updated_at, last_login FROM "user" WHERE id = ?`

	query = sqlDb.Rebind(query)
	row := sqlDb.QueryRow(query, id)

	var u userWithPassword
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.LastLogin); err != nil {
		t.Fatal(err)
	}

	return &u
}
