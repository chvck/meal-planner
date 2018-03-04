package server

import "github.com/chvck/meal-planner/controller"

// Handler is a container for controllers, providing a single point of access
// for the router
type Handler struct {
	controller.MenuController
	controller.RecipeController
	controller.UserController
}

// NewHandler creates a new Handler
func NewHandler(mc controller.MenuController, rc controller.RecipeController, uc controller.UserController) *Handler {
	return &Handler{mc, rc, uc}
}
