package server

import "github.com/chvck/meal-planner/controller"

// Handler is a container for controllers, providing a single point of access
// for the router
type Handler struct {
	controller.Controller
}

// NewHandler creates a new Handler
func NewHandler(c controller.Controller) *Handler {
	return &Handler{c}
}
