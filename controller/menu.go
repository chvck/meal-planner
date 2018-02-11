package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chvck/meal-planner/model/menu"
	"github.com/chvck/meal-planner/model/user"
	"github.com/chvck/meal-planner/store"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// MenuByID is the HTTP handler for fetching a single menu
func MenuByID(w http.ResponseWriter, r *http.Request) {
	db := store.Database()
	vars := mux.Vars(r)
	u, ok := context.Get(r, "user").(user.User)
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

	menu, err := menu.One(db, id, u.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve menu", http.StatusNotFound)
		return
	}

	JSONResponse(menu, w)
}
