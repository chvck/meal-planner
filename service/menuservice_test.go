package service_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/guregu/null.v3"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
)

type TestMenuDataModel struct {
	OnForPlannersCalled func(ids ...interface{}) (map[int][]model.Menu, error)
	Menus               []model.Menu
}

func (tmd TestMenuDataModel) One(id int, userID int) (*model.Menu, error) {
	for _, m := range tmd.Menus {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, nil
}

func (tmd TestMenuDataModel) AllWithLimit(limit int, offset int, userID int) ([]model.Menu, error) {
	return tmd.Menus[offset : limit+offset], nil
}

// Defer the result of this to the caller
func (tmd TestMenuDataModel) ForPlanners(ids ...interface{}) (map[int][]model.Menu, error) {
	return tmd.OnForPlannersCalled(ids...)
}

func (tmd TestMenuDataModel) Create(m model.Menu, userID int) (*int, error) {
	id := 10
	return &id, nil
}

func (tmd TestMenuDataModel) Update(m model.Menu, id int, userID int) error {
	return nil
}

func (tmd TestMenuDataModel) Delete(id int, userID int) error {
	return nil
}

func NewTestMenuDataModel(menus []model.Menu) *TestMenuDataModel {
	menuDataModel := TestMenuDataModel{
		Menus: menus,
	}

	return &menuDataModel
}

func TestMenuServiceGetByID(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(1))
	expected := menus(1)[0]

	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	m, err := service.GetByID(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func TestMenuServiceGetByIDWithRecipes(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(1))
	expected := menus(1)[0]

	recipes := recipesForMenus(1)

	rDataModel := TestRecipeDataModel{}
	rDataModel.OnForMenusCalled = func(ids ...interface{}) (map[int][]model.Recipe, error) {
		actual := make(map[int][]model.Recipe)
		actual[1] = recipes[1]
		return actual, nil
	}

	expected.Recipes = recipes[1]

	service := service.NewMenuService(mDataModel, rDataModel)

	m, err := service.GetByIDWithRecipes(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func TestMenuServiceAll(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(20))
	menus := menus(20)

	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	result, err := service.All(10, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, menus[0:10], result)
}

func TestMenuServiceAllLimit(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(20))
	menus := menus(20)

	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	result, err := service.All(5, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, menus[0:5], result)
}

func TestMenuServiceAllOffset(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(20))
	menus := menus(20)

	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	result, err := service.All(10, 5, 1)

	assert.Nil(t, err)
	assert.Equal(t, menus[5:15], result)
}

func TestMenuServiceAllWithRecipes(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(20))
	menus := menus(20)

	rDataModel := TestRecipeDataModel{}
	recipes := recipesForMenus(20)

	rDataModel.OnForMenusCalled = func(ids ...interface{}) (map[int][]model.Recipe, error) {
		actual := recipes
		return actual, nil
	}

	service := service.NewMenuService(mDataModel, rDataModel)

	result, err := service.AllWithRecipes(10, 0, 1)
	expected := menus[0:10]
	for i := range expected {
		expected[i].Recipes = recipes[expected[i].ID]
	}

	assert.Nil(t, err)
	assert.Equal(t, menus[0:10], result)
}

func TestMenuServiceCreate(t *testing.T) {
	mDataModel := NewTestMenuDataModel(menus(20))
	rDataModel := TestRecipeDataModel{}

	service := service.NewMenuService(mDataModel, rDataModel)

	m := model.Menu{
		Name:        "test0",
		Description: null.StringFrom("desc0"),
		UserID:      1,
	}
	result, err := service.Create(m, 1)

	expected := model.Menu{
		Name:        "test9",
		Description: null.StringFrom("desc9"),
		UserID:      1,
		ID:          10,
	}

	assert.Nil(t, err)
	assert.Equal(t, &expected, result)
}

func menus(numMenus int) []model.Menu {
	menus := []model.Menu{}

	for i := 0; i < numMenus; i++ {
		m := model.Menu{
			ID:          i + 1,
			Name:        fmt.Sprintf("test%v", i),
			Description: null.StringFrom(fmt.Sprintf("desc%v", i)),
			UserID:      1,
		}

		menus = append(menus, m)
	}

	return menus
}

func recipesForMenus(numRecipes int) map[int][]model.Recipe {
	recipes := make(map[int][]model.Recipe)
	for i := 0; i < numRecipes; i++ {
		recipes[i] = []model.Recipe{
			model.Recipe{
				ID:           i,
				Name:         fmt.Sprintf("recipe %v", i),
				Instructions: fmt.Sprintf("instructions %v", i),
				Description:  null.StringFrom(fmt.Sprintf("description %v", i)),
				CookTime:     null.IntFrom(15),
				PrepTime:     null.IntFrom(25),
				UserID:       1,
				Yield:        null.IntFrom(2),
			},
		}
	}

	return recipes
}
