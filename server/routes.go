package server

import (
	"github.com/gorilla/mux"
	"github.com/chvck/meal-planner/controller"
)

func routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/recipe/", controller.RecipeIndex)

	return router
}
