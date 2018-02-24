package server

import "github.com/chvck/meal-planner/controller"

type Handler struct {
	controller.MenuController
	controller.RecipeController
	controller.UserController
}

func NewHandler(mc controller.MenuController, rc controller.RecipeController, uc controller.UserController) *Handler {
	return &Handler{mc, rc, uc}
}
