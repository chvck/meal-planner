package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// RecipeController is the interface for a controller than handles recipe endpoints
type RecipeController interface {
	RecipeIndex(w http.ResponseWriter, r *http.Request)
	RecipeByID(w http.ResponseWriter, r *http.Request)
	RecipeCreate(w http.ResponseWriter, r *http.Request)
	RecipeUpdate(w http.ResponseWriter, r *http.Request)
	RecipeDelete(w http.ResponseWriter, r *http.Request)
}

type recipeController struct {
	service service.RecipeService
}

// NewRecipeController creates a new recipe controller
func NewRecipeController(service service.RecipeService) RecipeController {
	return &recipeController{service: service}
}

const (
	defaultRecipePerPage = 10
	defaultRecipeOffset  = 0
)

// RecipeIndex is the HTTP handler for the recipe index endpoint
func (rc recipeController) RecipeIndex(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	perPage := getURLParameterAsInt(r.URL, "perPage", defaultRecipePerPage)
	offset := getURLParameterAsInt(r.URL, "offset", defaultRecipeOffset)
	recipes, err := rc.service.All(perPage, offset, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(recipes, w)
}

// RecipeByID is the HTTP handler for fetching a single recipe
func (rc recipeController) RecipeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u := context.Get(r, "user").(model.User)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	recipe, err := rc.service.GetByIDWithIngredients(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(recipe, w)
}

// RecipeCreate is the HTTP handler for creating a recipe
func (rc recipeController) RecipeCreate(w http.ResponseWriter, r *http.Request) {
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
	created, err := rc.service.Create(re, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(created, w, 201)
}

// RecipeUpdate is the HTTP handler for updating a recipe
func (rc recipeController) RecipeUpdate(w http.ResponseWriter, r *http.Request) {
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

	updated, err := rc.service.Update(re, id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(updated, w, 200)
}

// RecipeDelete is the HTTP handler for deleting a recipe
func (rc recipeController) RecipeDelete(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	err = rc.service.Delete(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
}
