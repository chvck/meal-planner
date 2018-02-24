package service

import (
	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

type MenuService struct {
	mdm datamodel.MenuDataModel
	rdm datamodel.RecipeDataModel
}

// GetByID retrieves a menu by id
func (ms MenuService) GetByID(id int, userID int) (*model.Menu, error) {
	return ms.mdm.One(id, userID)
}

// GetByIDWithRecipes retrieves a menu by id, including its recipes
func (ms MenuService) GetByIDWithRecipes(id int, userID int) (*model.Menu, error) {
	m, err := ms.mdm.One(id, userID)
	if err != nil {
		return nil, err
	}

	recipes, err := ms.rdm.ForMenus(m.ID)
	if err != nil {
		return nil, err
	}
	if recipes[m.ID] != nil {
		m.Recipes = recipes[m.ID]
	}

	return m, nil
}

// All retrieves all menus
func (ms MenuService) All(limit int, offset int, userID int) ([]model.Menu, error) {
	return ms.mdm.AllWithLimit(limit, offset, userID)
}

// AllWithRecipes retrieves all menus, with recipes
func (ms MenuService) AllWithRecipes(limit int, offset int, userID int) ([]model.Menu, error) {
	menus, err := ms.mdm.AllWithLimit(limit, offset, userID)
	if err != nil {
		return nil, err
	}

	menuIDs := make([]interface{}, len(menus))
	for i, m := range menus {
		menuIDs[i] = m.ID
	}

	// if there aren't any menus then return empty slice
	if len(menus) == 0 {
		return menus, nil
	}

	recipesByMenuID, err := ms.rdm.ForMenus(menuIDs...)
	if err != nil {
		return nil, err
	}

	for i, m := range menus {
		recipes, ok := recipesByMenuID[m.ID]
		if ok {
			menus[i].Recipes = recipes
		}
	}

	return menus, nil
}

// Create creates a new menu
func (ms MenuService) Create(m model.Menu, userID int) (*model.Menu, error) {
	mID, err := ms.mdm.Create(m, userID)
	if err != nil {
		return nil, err
	}

	return ms.mdm.One(*mID, userID)
}

// Update updates a menu
func (ms MenuService) Update(m model.Menu, id int, userID int) error {
	return ms.mdm.Update(m, id, userID)
}

// Delete deletes a menu
func (ms MenuService) Delete(id int, userID int) error {
	return ms.mdm.Delete(id, userID)
}
