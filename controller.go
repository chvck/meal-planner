package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Controller interface {
	RecipeIndex(w http.ResponseWriter, r *http.Request)
	RecipeByID(w http.ResponseWriter, r *http.Request)
	RecipeCreate(w http.ResponseWriter, r *http.Request)
	RecipeUpdate(w http.ResponseWriter, r *http.Request)
	RecipeDelete(w http.ResponseWriter, r *http.Request)
	UserLogin(w http.ResponseWriter, r *http.Request)
	UserCreate(w http.ResponseWriter, r *http.Request)
}

type StandardController struct {
	ds      DataStore
	authKey string
}

func NewStandardController(ds DataStore, authKey string) *StandardController {
	return &StandardController{
		ds:      ds,
		authKey: authKey,
	}
}

func getURLParameterAsInt(rURL *url.URL, param string, defaultVal int) int {
	query := rURL.Query()
	val, _ := strconv.Atoi(query.Get(param))
	if val == 0 {
		val = defaultVal
	}

	return val
}

const (
	defaultRecipePerPage = 10
	defaultRecipeOffset  = 0
)

// RecipeIndex is the HTTP handler for the recipe index endpoint
func (sc StandardController) RecipeIndex(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(User)
	perPage := getURLParameterAsInt(r.URL, "perPage", defaultRecipePerPage)
	offset := getURLParameterAsInt(r.URL, "offset", defaultRecipeOffset)
	recipes, err := sc.ds.Recipes(perPage, offset, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(recipes, w)
}

// RecipeByID is the HTTP handler for fetching a single recipe
func (sc StandardController) RecipeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u := context.Get(r, "user").(User)

	recipe, err := sc.ds.Recipe(vars["id"], u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(recipe, w)
}

// RecipeCreate is the HTTP handler for creating a recipe
func (sc StandardController) RecipeCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	var re Recipe
	if err := json.Unmarshal(body, &re); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	errs := re.Validate()
	if len(errs) != 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	u := context.Get(r, "user").(User)
	created, err := sc.ds.RecipeCreate(re, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(created, w, 201)
}

// RecipeUpdate is the HTTP handler for updating a recipe
func (sc StandardController) RecipeUpdate(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(User)
	vars := mux.Vars(r)

	var re Recipe
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &re); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	errs := re.Validate()
	if len(errs) != 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	err = sc.ds.RecipeUpdate(re, vars["id"], u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(re, w, 200)
}

// RecipeDelete is the HTTP handler for deleting a recipe
func (sc StandardController) RecipeDelete(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(User)
	vars := mux.Vars(r)

	err := sc.ds.RecipeDelete(vars["id"], u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}

type userWithPassword struct {
	User
	Password string `json:"password,omitzero"`
}

type loginCredentials struct {
	Username string `json:"username,omitzero"`
	Password string `json:"password,omitzero"`
}

type jwtToken struct {
	Token string `json:"token,omitzero"`
}

// UserLogin is the HTTP handler for logging as user into the system
func (sc StandardController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	u := sc.ds.UserValidatePassword(creds.Username, []byte(creds.Password))
	if u == nil {
		err := errors.New("invalid user credentials provided")
		JSONResponseWithCode(NewJSONError(err), w, http.StatusUnauthorized)
		return
	}

	t, err := createToken(u, sc.authKey)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(t, w)
}

// UserCreate is the HTTP handler for creating a user
func (sc StandardController) UserCreate(w http.ResponseWriter, r *http.Request) {
	var u userWithPassword

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &u); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	if errs := u.Validate(); len(errs) > 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	if err := ValidatePassword(u.Password); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	created, err := sc.ds.UserCreate(u.User, []byte(u.Password))
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(created, w, 201)
}

func createToken(user *User, key string) (*jwtToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.ID,
		"lastLogin": user.LastLogin,
		"level":     LevelUser,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}
	return &jwtToken{Token: tokenString}, err
}

// JSONError is a wrapper for error/errors that can be parsed into JSON
type JSONError struct {
	Errors []string `json:"errors"`
}

// NewJSONError creates a new JSONError from a single error
func NewJSONError(err error) JSONError {
	errs := []error{err}
	return NewJSONErrors(errs)
}

// NewJSONErrors creates a new JSONError from a collection of errors
func NewJSONErrors(errs []error) JSONError {
	jsonErr := JSONError{}
	for _, err := range errs {
		jsonErr.Errors = append(jsonErr.Errors, err.Error())
	}

	return jsonErr
}

// JSONResponse parses the data into JSON and writes it into the response
func JSONResponse(response interface{}, w http.ResponseWriter) {
	JSONResponseWithCode(response, w, 200)
}

// JSONResponseWithCode parses the data into JSON and writes it into the response
// returning response with status code supplied
func JSONResponseWithCode(response interface{}, w http.ResponseWriter, statusCode int) {
	js, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

