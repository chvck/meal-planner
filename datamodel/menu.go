package datamodel

import "github.com/chvck/meal-planner/model"

// MenuDataModel is the datamodel for data store operations on the Menu model
type MenuDataModel interface {
	One(id int, userID int) (*model.Menu, error)
	AllWithLimit(limit int, offset int, userID int) ([]model.Menu, error)
	ForPlanners(ids ...interface{}) (map[int][]model.Menu, error)
	Create(m model.Menu, userID int) (*int, error)
	Update(m model.Menu, id int, userID int) error
	Delete(id int, userID int) error
	AddRecipe(menuID int, recipeID int) error
	RemoveRecipe(menuID int, recipeID int) error
}
