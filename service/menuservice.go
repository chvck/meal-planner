package service

import (
	"errors"

	"github.com/chvck/meal-planner/datamodel"
	"github.com/chvck/meal-planner/model"
)

// MenuService is the service for interacting with Menus
type MenuService interface {
	GetByID(id int, userID int) (*model.Menu, error)
	GetByIDWithRecipes(id int, userID int) (*model.Menu, error)
	All(limit int, offset int, userID int) ([]model.Menu, error)
	AllWithRecipes(limit int, offset int, userID int) ([]model.Menu, error)
	Create(m model.Menu, userID int) (*model.Menu, error)
	Update(m model.Menu, id int, userID int) (*model.Menu, error)
	Delete(id int, userID int) error
	AddRecipe(menuID int, recipeID int, userID int) error
	RemoveRecipe(menuID int, recipeID int, userID int) error
}

type menuService struct {
	mdm datamodel.MenuDataModel
	rdm datamodel.RecipeDataModel
}

// NewMenuService creates a new menu service
func NewMenuService(mdm datamodel.MenuDataModel, rdm datamodel.RecipeDataModel) MenuService {
	return &menuService{mdm: mdm, rdm: rdm}
}

// GetByID retrieves a menu by id
func (ms menuService) GetByID(id int, userID int) (*model.Menu, error) {
	return ms.mdm.One(id, userID)
}

// GetByIDWithRecipes retrieves a menu by id, including its recipes
func (ms menuService) GetByIDWithRecipes(id int, userID int) (*model.Menu, error) {
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
func (ms menuService) All(limit int, offset int, userID int) ([]model.Menu, error) {
	return ms.mdm.AllWithLimit(limit, offset, userID)
}

// AllWithRecipes retrieves all menus, with recipes
func (ms menuService) AllWithRecipes(limit int, offset int, userID int) ([]model.Menu, error) {
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
		} else {
			// make sure that this an empty array rather than a nil pointer
			menus[i].Recipes = []model.Recipe{}
		}
	}

	return menus, nil
}

// Create creates a new menu
func (ms menuService) Create(m model.Menu, userID int) (*model.Menu, error) {
	mID, err := ms.mdm.Create(m, userID)
	if err != nil {
		return nil, err
	}

	return ms.mdm.One(*mID, userID)
}

// Update updates a menu
func (ms menuService) Update(m model.Menu, id int, userID int) (*model.Menu, error) {
	err := ms.mdm.Update(m, id, userID)
	if err != nil {
		return nil, err
	}

	return ms.mdm.One(id, userID)
}

// Delete deletes a menu
func (ms menuService) Delete(id int, userID int) error {
	return ms.mdm.Delete(id, userID)
}

// AddRecipe adds a recipe to a menu
func (ms menuService) AddRecipe(menuID int, recipeID int, userID int) error {
	m, err := ms.mdm.One(menuID, userID)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.New("menu could not be found")
	}

	r, err := ms.rdm.One(recipeID, userID)
	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("recipe could not be found")
	}

	return ms.mdm.AddRecipe(menuID, recipeID)
}

// RemoveRecipe removes a recipe from a menu
func (ms menuService) RemoveRecipe(menuID int, recipeID int, userID int) error {
	m, err := ms.mdm.One(menuID, userID)
	if err != nil {
		return err
	}

	if m == nil {
		return errors.New("menu could not be found")
	}

	r, err := ms.rdm.One(recipeID, userID)
	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("recipe could not be found")
	}

	return ms.mdm.RemoveRecipe(menuID, recipeID)
}
