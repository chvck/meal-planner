package controller

import (
	"net/http"
	"github.com/chvck/meal-planner/context"
	"github.com/chvck/meal-planner/model/recipe"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"
)

func RecipeIndex(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	recipes, err := recipe.All(db)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not retrieve recipes", http.StatusNotFound)
		return
	}

	JsonResponse(recipes, w)
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

	JsonResponse(recipes, w)
}

func RecipeCreate(w http.ResponseWriter, r *http.Request) {
	db := context.Database()
	var re recipe.Recipe
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid recipe", http.StatusBadRequest)
		return
	} else {
		if err := json.Unmarshal(body, &re); err != nil {
			log.Println(err.Error())
			http.Error(w, "Invalid recipe", http.StatusBadRequest)
			return
		}
	}

	err := recipe.Create(db, re)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Could not create recipe", http.StatusInternalServerError)
		return
	}

}
