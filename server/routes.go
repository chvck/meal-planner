package server

import (
	"github.com/gorilla/mux"
	"github.com/chvck/meal-planner/controller"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/recipe/", controller.RecipeIndex)
	router.HandleFunc("/recipe/{id}", controller.RecipeById)

	router.HandleFunc("/ingredient/", controller.IngredientIndex)

	return router
}
