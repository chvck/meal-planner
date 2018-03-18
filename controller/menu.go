package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// MenuController is the interface for a controller than handles menu endpoints
type MenuController interface {
	MenuIndex(w http.ResponseWriter, r *http.Request)
	MenuByID(w http.ResponseWriter, r *http.Request)
	MenuCreate(w http.ResponseWriter, r *http.Request)
	MenuUpdate(w http.ResponseWriter, r *http.Request)
	MenuDelete(w http.ResponseWriter, r *http.Request)
	MenuAddRecipe(w http.ResponseWriter, r *http.Request)
	MenuRemoveRecipe(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	service service.MenuService
}

// NewMenuController creates a new recipe controller
func NewMenuController(service service.MenuService) MenuController {
	return &menuController{service: service}
}

const (
	defaultMenuPerPage = 10
	defaultMenuOffset  = 0
)

// MenuIndex is the HTTP handler for the menu index endpoint
func (mc menuController) MenuIndex(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	perPage := getURLParameterAsInt(r.URL, "perPage", defaultMenuPerPage)
	offset := getURLParameterAsInt(r.URL, "offset", defaultMenuOffset)
	q := r.URL.Query()
	includeRecipes := q.Get("includeRecipes")
	var menuFunc func(int, int, int) ([]model.Menu, error)
	if includeRecipes == "true" {
		menuFunc = mc.service.AllWithRecipes
	} else {
		menuFunc = mc.service.All
	}

	menus, err := menuFunc(perPage, offset, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(menus, w)
}

// MenuByID is the HTTP handler for fetching a single menu
func (mc menuController) MenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, ok := context.Get(r, "user").(model.User)
	if !ok {
		err := errors.New("Cannot extract user from request")
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	includeRecipes := q.Get("includeRecipes")
	var menuFunc func(int, int) (*model.Menu, error)
	if includeRecipes == "true" {
		menuFunc = mc.service.GetByIDWithRecipes
	} else {
		menuFunc = mc.service.GetByID
	}

	menu, err := menuFunc(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponse(menu, w)
}

// MenuCreate is the HTTP handler for creating a menu
func (mc menuController) MenuCreate(w http.ResponseWriter, r *http.Request) {
	var m model.Menu
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	errs := m.Validate()
	if len(errs) != 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	u := context.Get(r, "user").(model.User)
	created, err := mc.service.Create(m, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(created, w, 201)
}

// RecipeUpdate is the HTTP handler for updating a recipe
func (mc menuController) MenuUpdate(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	var m model.Menu
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusBadRequest)
		return
	}

	errs := m.Validate()
	if len(errs) != 0 {
		JSONResponseWithCode(NewJSONErrors(errs), w, http.StatusBadRequest)
		return
	}

	updated, err := mc.service.Update(m, id, u.ID)
	if err != nil {
		log.Println(err.Error())
		JSONResponseWithCode(NewJSONError(err), w, http.StatusInternalServerError)
		return
	}

	JSONResponseWithCode(updated, w, 200)
}

// MenuDelete is the HTTP handler for deleting a menu
func (mc menuController) MenuDelete(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = mc.service.Delete(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not delete menu", http.StatusInternalServerError)
		return
	}
}

// MenuAddRecipe is the HTTP handler for adding a recipe to a menu
func (mc menuController) MenuAddRecipe(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)

	mID, err := strconv.Atoi(vars["mId"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
		return
	}

	rID, err := strconv.Atoi(vars["rId"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	err = mc.service.AddRecipe(mID, rID, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not add recipe to menu", http.StatusInternalServerError)
		return
	}
}

// MenuRemoveRecipe is the HTTP handler for removing a recipe from a menu
func (mc menuController) MenuRemoveRecipe(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)

	mID, err := strconv.Atoi(vars["mId"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
		return
	}

	rID, err := strconv.Atoi(vars["rId"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid recipe ID", http.StatusBadRequest)
		return
	}

	err = mc.service.RemoveRecipe(mID, rID, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not add recipe to menu", http.StatusInternalServerError)
		return
	}
}
