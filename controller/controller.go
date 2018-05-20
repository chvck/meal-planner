package controller

import (
	"net/url"
	"strconv"
	"net/http"

	"github.com/chvck/meal-planner/datastore"
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
	ds datastore.DataStore
	authKey string
}

func NewStandardController(ds datastore.DataStore, authKey string) *StandardController {
	return &StandardController{
		ds: ds,
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
