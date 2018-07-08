package datastore

import "github.com/chvck/meal-planner/model"

// DataStore is used to access data from the underlying store
type DataStore interface {
	User(id string) (*model.User, error)
	Users(limit, offset int) ([]model.User, error)
	UserCreate(u model.User, password []byte) (*model.User, error)
	UserValidatePassword(username string, pw []byte) *model.User
	Recipe(id, userID string) (*model.Recipe, error)
	Recipes(limit, offset int, userID string) ([]model.Recipe, error)
	RecipeCreate(r model.Recipe, userID string) (*model.Recipe, error)
	RecipeUpdate(r model.Recipe, id, userID string) error
	RecipeDelete(id, userID string) error
	PlannerWithRecipeNames(id, userID string) (*model.Planner, error)
	PlannersWithRecipeNames(start, end int, userID string) ([]model.Planner, error)
	PlannerCreate(when int, mealtime, userID string) (*model.Planner, error)
	PlannerAddRecipe(plannerID, recipeID, userID string) error
	PlannerRemoveRecipe(plannerID, recipeID, userID string) error
}

// ConnectionConfig holds configuration details for connecting to a DataStore
type ConnectionConfig struct {
	server   string
	port     int
	username string
	password string
}
