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

// MenuController is the interface for a controller than handles menu endpoints
type MenuController interface {
	MenuIndex(w http.ResponseWriter, r *http.Request)
	MenuByID(w http.ResponseWriter, r *http.Request)
	MenuCreate(w http.ResponseWriter, r *http.Request)
	MenuUpdate(w http.ResponseWriter, r *http.Request)
	MenuDelete(w http.ResponseWriter, r *http.Request)
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
	menus, err := mc.service.All(perPage, offset, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve menus", http.StatusNotFound)
		return
	}

	JSONResponse(menus, w)
}

// MenuByID is the HTTP handler for fetching a single menu
func (mc menuController) MenuByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	u, ok := context.Get(r, "user").(model.User)
	if !ok {
		log.Println("Cannot extract user from request")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	menu, err := mc.service.GetByID(id, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve menu", http.StatusNotFound)
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
		http.Error(w, "Invalid menu", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid menu", http.StatusBadRequest)
		return
	}

	u := context.Get(r, "user").(model.User)
	created, err := mc.service.Create(m, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create menu", http.StatusInternalServerError)
		return
	}

	JSONResponse(created, w)
}

// RecipeUpdate is the HTTP handler for updating a recipe
func (mc menuController) MenuUpdate(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var m model.Menu
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid menu", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid menu", http.StatusBadRequest)
		return
	}

	err = mc.service.Update(m, id, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create menu", http.StatusInternalServerError)
		return
	}
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
