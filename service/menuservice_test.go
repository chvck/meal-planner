package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/guregu/null.v3"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
)

type TestMenuDataModel struct {
	OnOneCalled         func(id int, userID int) (*model.Menu, error)
	OnAllCalled         func(limit int, offset int, userID int) ([]model.Menu, error)
	OnForPlannersCalled func(ids ...interface{}) (map[int][]model.Menu, error)
	OnCreateCalled      func(m model.Menu, userID int) (*int, error)
}

func (tmd TestMenuDataModel) One(id int, userID int) (*model.Menu, error) {
	return tmd.OnOneCalled(id, userID)
}

func (tmd TestMenuDataModel) AllWithLimit(limit int, offset int, userID int) ([]model.Menu, error) {
	return tmd.OnAllCalled(limit, offset, userID)
}

// Defer the result of this to the caller
func (tmd TestMenuDataModel) ForPlanners(ids ...interface{}) (map[int][]model.Menu, error) {
	return tmd.OnForPlannersCalled(ids...)
}

func (tmd TestMenuDataModel) Create(m model.Menu, userID int) (*int, error) {
	return tmd.OnCreateCalled(m, userID)
}

func (tmd TestMenuDataModel) Update(m model.Menu, id int, userID int) error {
	return nil
}

func (tmd TestMenuDataModel) Delete(id int, userID int) error {
	return nil
}

func TestMenuServiceGetByID(t *testing.T) {
	mDataModel := TestMenuDataModel{}
	expected := model.Menu{
		ID:          1,
		Name:        "Test",
		Description: null.StringFrom("Test desc"),
		UserID:      1,
	}
	mDataModel.OnOneCalled = func(id int, userID int) (*model.Menu, error) {
		return &expected, nil
	}

	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	m, err := service.GetByID(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func TestMenuServiceGetByIDWithRecipes(t *testing.T) {
	mDataModel := TestMenuDataModel{}
	expected := model.Menu{
		ID:          1,
		Name:        "Test",
		Description: null.StringFrom("Test desc"),
		UserID:      1,
	}
	mDataModel.OnOneCalled = func(id int, userID int) (*model.Menu, error) {
		actual := expected
		return &actual, nil
	}

	recipes := recipes()

	rDataModel := TestRecipeDataModel{}
	rDataModel.OnForMenusCalled = func(ids ...interface{}) (map[int][]model.Recipe, error) {
		return recipes, nil
	}

	expected.Recipes = recipes[1]

	service := service.NewMenuService(mDataModel, rDataModel)

	m, err := service.GetByIDWithRecipes(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func TestMenuServiceAll(t *testing.T) {
	mDataModel := TestMenuDataModel{}
	expected := model.Menu{
		ID:          1,
		Name:        "Test",
		Description: null.StringFrom("Test desc"),
		UserID:      1,
	}
	mDataModel.OnAllCalled = func(id int, userID int) (*model.Menu, error) {
		actual := expected
		return &actual, nil
	}

	recipes := recipes()

	rDataModel := TestRecipeDataModel{}
	rDataModel.OnForMenusCalled = func(ids ...interface{}) (map[int][]model.Recipe, error) {
		return recipes, nil
	}

	expected.Recipes = recipes[1]

	service := service.NewMenuService(mDataModel, rDataModel)

	m, err := service.GetByIDWithRecipes(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func recipes() map[int][]model.Recipe {
	recipes := make(map[int][]model.Recipe)
	recipes[1] = []model.Recipe{
		model.Recipe{
			ID:           1,
			Name:         "recipe 1",
			Instructions: "instructions 1",
			Description:  null.StringFrom("description 1"),
			CookTime:     null.IntFrom(15),
			PrepTime:     null.IntFrom(25),
			UserID:       1,
			Yield:        null.IntFrom(2),
		},
		model.Recipe{
			ID:           2,
			Name:         "recipe 2",
			Instructions: "instructions 2",
			Description:  null.StringFrom("description 2"),
			CookTime:     null.IntFrom(12),
			PrepTime:     null.IntFrom(22),
			UserID:       1,
			Yield:        null.IntFrom(2),
		},
	}
	recipes[2] = []model.Recipe{
		model.Recipe{
			ID:           3,
			Name:         "recipe 3",
			Instructions: "instructions 3",
			Description:  null.StringFrom("description 3"),
			CookTime:     null.IntFrom(10),
			PrepTime:     null.IntFrom(20),
			UserID:       1,
			Yield:        null.IntFrom(1),
		},
	}

	return recipes
}
