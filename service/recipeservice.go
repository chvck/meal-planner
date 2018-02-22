package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model/recipe"
)

type RecipeService struct {
	dm datamodel.RecipeDataModel
}

// GetByIDWithIngredients retrieves a recipe by id
func (rs RecipeService) GetByIDWithIngredients(id int, userID int) (*recipe.Recipe, error) {
	return rs.dm.One(id, userID)
}

// All retrieves all recipes
func (rs RecipeService) All(limit int, offset int, userID int) ([]recipe.Recipe, error) {
	return rs.dm.AllWithLimit(limit, offset, userID)
}

// Create creates a new recipe
func (rs RecipeService) Create(r recipe.Recipe, userID int) (*recipe.Recipe, error) {
	rID, err := rs.dm.Create(r, userID)
	if err != nil {
		return nil, err
	}

	return rs.dm.One(*rID, userID)
}

// Update updates a recipe
func (rs RecipeService) Update(r recipe.Recipe, id int, userID int) error {
	return rs.dm.Update(r, id, userID)
}

// Delete delete a recipe
func (rs RecipeService) Delete(id int, userID int) error {
	return rs.dm.Delete(id, userID)
}
