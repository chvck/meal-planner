package datamodel

import "github.com/chvck/meal-planner/model/planner"

// PlannerDataModel is the datamodel for data store operations on the Planner model
type PlannerDataModel interface {
	All(int, int, int) ([]planner.Planner, error)
	One(int, string, int) (*planner.Planner, error)
	Create(int, string, int) (*int, error)
	AddRecipe(int, int) error
	RemoveRecipe(int, string, int, int) error
}
