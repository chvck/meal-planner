package main

// Handler is a container for controllers, providing a single point of access
// for the router
type Handler struct {
	Controller
}

// NewHandler creates a new Handler
func NewHandler(c Controller) *Handler {
	return &Handler{c}
}

