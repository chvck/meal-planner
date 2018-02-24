package datamodel

import "github.com/chvck/meal-planner/model"

// RecipeDataModel is the datamodel for data store operations on the Recipe model
type RecipeDataModel interface {
	FindByIngredientNames(names ...interface{}) ([]model.Recipe, error)
	One(id int, userID int) (*model.Recipe, error)
	AllWithLimit(limit int, offset int, userID int) ([]model.Recipe, error)
	ForMenus(ids ...interface{}) (map[int][]model.Recipe, error)
	ForPlanners(ids ...interface{}) (map[int][]model.Recipe, error)
	Create(r model.Recipe, userID int) (*int, error)
	Update(r model.Recipe, id int, userID int) error
	Delete(id int, userID int) error
}
