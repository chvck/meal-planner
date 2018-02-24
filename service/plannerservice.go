package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

type PlannerService struct {
	mdm datamodel.MenuDataModel
	rdm datamodel.RecipeDataModel
	pdm datamodel.PlannerDataModel
}

// All retrieves all planners
func (ps PlannerService) All(limit int, offset int, userID int) ([]model.Planner, error) {
	return ps.pdm.All(limit, offset, userID)
}

// AllWithRelations retrieves all planners, with menus and recipes
func (ps PlannerService) AllWithRelations(limit int, offset int, userID int) ([]model.Planner, error) {
	planners, err := ps.pdm.All(limit, offset, userID)
	if err != nil {
		return nil, err
	}

	// if there aren't any planners then return empty slice
	if len(planners) == 0 {
		return planners, nil
	}

	plannerIDs := make([]interface{}, len(planners))
	for i, p := range planners {
		plannerIDs[i] = p.ID
	}

	recipesByPlannerID, err := ps.rdm.ForPlanners(plannerIDs...)
	if err != nil {
		return nil, err
	}

	// add menus

	for i, p := range planners {
		recipes, ok := recipesByPlannerID[p.ID]
		if ok {
			planners[i].Recipes = recipes
		}
		// menus, ok := menusByPlannerID[p.ID]
		// if ok {
		// 	planners[i].Menus = menus
		// }
	}

	return planners, nil
}

// AddRecipe adds a recipe to a planner, will create the planner if it doesn't exist
func (ps PlannerService) AddRecipe(when int, mealtime string, recipeID int, userID int) error {
	p, err := ps.pdm.One(when, mealtime, userID)
	if err != nil {
		return err
	}

	var pID int
	if p == nil {
		id, err := ps.pdm.Create(when, mealtime, userID)
		if err != nil {
			return err
		}
		pID = *id
	} else {
		pID = p.ID
	}

	return ps.pdm.AddRecipe(pID, recipeID)
}

// RemoveRecipe removes a recipe from a planner
func (ps PlannerService) RemoveRecipe(when int, mealtime string, recipeID int, userID int) error {
	return ps.pdm.RemoveRecipe(when, mealtime, recipeID, userID)
}
