package service_test

import (
	"github.com/chvck/meal-planner/model"
)

type TestRecipeDataModel struct {
	Data                []model.Recipe
	OnForMenusCalled    func(ids ...interface{}) (map[int][]model.Recipe, error)
	OnForPlannersCalled func(ids ...interface{}) (map[int][]model.Recipe, error)
	IDToCreate          int
}

func (trd TestRecipeDataModel) FindByIngredientNames(names ...interface{}) ([]model.Recipe, error) {
	return nil, nil
}

func (trd TestRecipeDataModel) One(id int, userID int) (*model.Recipe, error) {
	for _, model := range trd.Data {
		if model.ID == id {
			return &model, nil
		}
	}
	return nil, nil
}

func (trd TestRecipeDataModel) AllWithLimit(limit int, offset int, userID int) ([]model.Recipe, error) {
	models := trd.Data[offset : offset+limit]
	return models, nil
}

// Defer the result of this to the caller
func (trd TestRecipeDataModel) ForMenus(ids ...interface{}) (map[int][]model.Recipe, error) {
	return trd.OnForMenusCalled(ids...)
}

// Defer the result of this to the caller
func (trd TestRecipeDataModel) ForPlanners(ids ...interface{}) (map[int][]model.Recipe, error) {
	return trd.OnForPlannersCalled(ids...)
}

func (trd TestRecipeDataModel) Create(m model.Recipe, userID int) (*int, error) {
	return &trd.IDToCreate, nil
}

func (trd TestRecipeDataModel) Update(m model.Recipe, id int, userID int) error {
	return nil
}

func (trd TestRecipeDataModel) Delete(id int, userID int) error {
	return nil
}
