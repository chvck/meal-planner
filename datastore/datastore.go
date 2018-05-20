package datastore

import "github.com/chvck/meal-planner/model"

// DataStore is used to access data from the underlying store
type DataStore interface {
	User(id int) (*model.User, error)
	Users(limit, offset int) ([]model.User, error)
	UserCreate(u model.User, password []byte) (*model.User, error)
	UserValidatePassword(username string, pw []byte) *model.User
	RecipesFromIngredientNames(names []string) ([]model.Recipe, error)
	Recipe(id, userID int) (*model.Recipe, error)
	Recipes(limit, offset, userID int) ([]model.Recipe, error)
	RecipeCreate(r model.Recipe, userID int) (*model.Recipe, error)
	RecipeUpdate(r model.Recipe, id, userID int) error
	RecipeDelete(id , userID int) error
	PlannerWithRecipeNames(when int, mealtime string, userID int) (*model.Planner, error)
	PlannersWithRecipeNames(start, end, userID int) ([]model.Planner, error)
	PlannerCreate(when int, mealtime string, userID int) (int, error)
	PlannerAddRecipe(plannerID, recipeID, userID int) error
	PlannerRemoveRecipe(plannerID, recipeID, userID int) error
}
