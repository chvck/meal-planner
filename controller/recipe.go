package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/chvck/meal-planner/model"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

const (
	defaultRecipePerPage = 10
	defaultRecipeOffset  = 0
)

// RecipeIndex is the HTTP handler for the recipe index endpoint
func (sc StandardController) RecipeIndex(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
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
	u := context.Get(r, "user").(model.User)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	recipe, err := sc.ds.Recipe(id, u.ID)
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

	var re model.Recipe
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

	u := context.Get(r, "user").(model.User)
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
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	var re model.Recipe
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

	err = sc.ds.RecipeUpdate(re, id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(re, w, 200)
}

// RecipeDelete is the HTTP handler for deleting a recipe
func (sc StandardController) RecipeDelete(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	err = sc.ds.RecipeDelete(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}
