package controller

import (
	"net/http"
	"fmt"
	"github.com/chvck/meal-planner/context"
	"github.com/chvck/meal-planner/model/recipe"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
)

func RecipeIndex(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	recipes, err := recipe.All(db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve recipes", http.StatusNotFound)
		return
	}

	js, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, "Could not retrieve recipes", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(js))
}

func RecipeById(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	recipes, err := recipe.One(db, id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve recipe", http.StatusNotFound)
		return
	}

	js, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, "Could not retrieve recipe", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(js))
}
