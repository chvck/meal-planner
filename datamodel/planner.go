package datamodel

import "github.com/chvck/meal-planner/model"

// PlannerDataModel is the datamodel for data store operations on the Planner model
type PlannerDataModel interface {
	All(int, int, int) ([]model.Planner, error)
	One(int, string, int) (*model.Planner, error)
	Create(int, string, int) (*int, error)
	AddRecipe(int, int) error
	RemoveRecipe(int, string, int, int) error
}
