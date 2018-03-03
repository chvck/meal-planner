package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

// RecipeService is the service for interacting with Recipes
type RecipeService interface {
	GetByIDWithIngredients(id int, userID int) (*model.Recipe, error)
	All(limit int, offset int, userID int) ([]model.Recipe, error)
	Create(r model.Recipe, userID int) (*model.Recipe, error)
	Update(r model.Recipe, id int, userID int) (*model.Recipe, error)
	Delete(id int, userID int) error
}

type recipeService struct {
	dm datamodel.RecipeDataModel
}

// NewRecipeService creates a new recipe service
func NewRecipeService(rdm datamodel.RecipeDataModel) RecipeService {
	return &recipeService{dm: rdm}
}

// GetByIDWithIngredients retrieves a recipe by id
func (rs recipeService) GetByIDWithIngredients(id int, userID int) (*model.Recipe, error) {
	return rs.dm.One(id, userID)
}

// All retrieves all recipes
func (rs recipeService) All(limit int, offset int, userID int) ([]model.Recipe, error) {
	return rs.dm.AllWithLimit(limit, offset, userID)
}

// Create creates a new recipe
func (rs recipeService) Create(r model.Recipe, userID int) (*model.Recipe, error) {
	rID, err := rs.dm.Create(r, userID)
	if err != nil {
		return nil, err
	}

	return rs.dm.One(*rID, userID)
}

// Update updates a recipe
func (rs recipeService) Update(r model.Recipe, id int, userID int) (*model.Recipe, error) {
	if err := rs.dm.Update(r, id, userID); err != nil {
		return nil, err
	}

	return rs.dm.One(id, userID)
}

// Delete delete a recipe
func (rs recipeService) Delete(id int, userID int) error {
	return rs.dm.Delete(id, userID)
}
