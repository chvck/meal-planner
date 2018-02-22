package datamodel

import (
	"github.com/chvck/meal-planner/model/recipe"
)

// RecipeDataModel is the datamodel for data store operations on the Recipe model
type RecipeDataModel interface {
	FindByIngredientNames(names ...interface{}) ([]recipe.Recipe, error)
	One(id int, userID int) (*recipe.Recipe, error)
	AllWithLimit(limit int, offset int, userID int) ([]recipe.Recipe, error)
	ForMenus(ids ...interface{}) (map[int][]recipe.Recipe, error)
	ForPlanners(ids ...interface{}) (map[int][]recipe.Recipe, error)
	Create(r recipe.Recipe, userID int) (*int, error)
	Update(r recipe.Recipe, id int, userID int) error
	Delete(id int, userID int) error
}
