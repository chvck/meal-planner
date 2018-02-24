package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type MenuController interface {
	MenuByID(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	service service.MenuService
}

func NewMenuController(service service.MenuService) MenuController {
	return &menuController{service: service}
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
