package service_test

import (
	"fmt"
	"testing"

	"github.com/chvck/meal-planner/model"
	"github.com/chvck/meal-planner/service"
	"github.com/stretchr/testify/assert"
	null "gopkg.in/guregu/null.v3"
)

type TestRecipeDataModel struct {
	Recipes             []model.Recipe
	OnForMenusCalled    func(ids ...interface{}) (map[int][]model.Recipe, error)
	OnForPlannersCalled func(ids ...interface{}) (map[int][]model.Recipe, error)
}

func (trd TestRecipeDataModel) FindByIngredientNames(names ...interface{}) ([]model.Recipe, error) {
	return nil, nil
}

func (trd TestRecipeDataModel) One(id int, userID int) (*model.Recipe, error) {
	for _, m := range trd.Recipes {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, nil
}

func (trd TestRecipeDataModel) AllWithLimit(limit int, offset int, userID int) ([]model.Recipe, error) {
	return trd.Recipes[offset : limit+offset], nil
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
	id := 10
	return &id, nil
}

func (trd TestRecipeDataModel) Update(m model.Recipe, id int, userID int) error {
	return nil
}

func (trd TestRecipeDataModel) Delete(id int, userID int) error {
	return nil
}

func NewTestRecipeDataModel(recipes []model.Recipe) *TestRecipeDataModel {
	recipeDataModel := TestRecipeDataModel{
		Recipes: recipes,
	}

	return &recipeDataModel
}

func TestRecipeServiceGetByIDWithIngredients(t *testing.T) {
	rDataModel := NewTestRecipeDataModel(recipes(1))
	expected := recipes(1)[0]

	service := service.NewRecipeService(rDataModel)

	m, err := service.GetByIDWithIngredients(1, 1)

	assert.Nil(t, err)
	assert.Equal(t, &expected, m)
}

func TestRecipeServiceAll(t *testing.T) {
	rDataModel := NewTestRecipeDataModel(recipes(20))
	expected := recipes(20)

	service := service.NewRecipeService(rDataModel)

	menus, err := service.All(10, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, expected[0:10], menus)
}

func TestRecipeServiceAllLimit(t *testing.T) {
	rDataModel := NewTestRecipeDataModel(recipes(20))
	expected := recipes(20)

	service := service.NewRecipeService(rDataModel)

	menus, err := service.All(2, 0, 1)

	assert.Nil(t, err)
	assert.Equal(t, expected[0:2], menus)
}

func TestRecipeServiceAllOffset(t *testing.T) {
	rDataModel := NewTestRecipeDataModel(recipes(20))
	expected := recipes(20)

	service := service.NewRecipeService(rDataModel)

	menus, err := service.All(10, 5, 1)

	assert.Nil(t, err)
	assert.Equal(t, expected[5:15], menus)
}

func TestRecipeServiceCreate(t *testing.T) {
	rDataModel := NewTestRecipeDataModel(recipes(20))

	service := service.NewRecipeService(rDataModel)

	r := model.Recipe{
		Name:         fmt.Sprintf("recipe 1"),
		Instructions: fmt.Sprintf("instructions 1"),
		Description:  null.StringFrom(fmt.Sprintf("description 1")),
		CookTime:     null.IntFrom(15),
		PrepTime:     null.IntFrom(25),
		UserID:       1,
		Yield:        null.IntFrom(2),
	}
	result, err := service.Create(r, 1)

	expected := model.Recipe{
		Name:         fmt.Sprintf("recipe 9"),
		Instructions: fmt.Sprintf("instructions 9"),
		Description:  null.StringFrom(fmt.Sprintf("description 9")),
		CookTime:     null.IntFrom(15),
		PrepTime:     null.IntFrom(25),
		UserID:       1,
		Yield:        null.IntFrom(2),
		ID:           10,
	}

	assert.Nil(t, err)
	assert.Equal(t, &expected, result)
}

func recipes(numRecipes int) []model.Recipe {
	recipes := *new([]model.Recipe)
	for i := 0; i < numRecipes; i++ {
		recipes = append(recipes, model.Recipe{
			ID:           i + 1,
			Name:         fmt.Sprintf("recipe %v", i),
			Instructions: fmt.Sprintf("instructions %v", i),
			Description:  null.StringFrom(fmt.Sprintf("description %v", i)),
			CookTime:     null.IntFrom(15),
			PrepTime:     null.IntFrom(25),
			UserID:       1,
			Yield:        null.IntFrom(2),
		},
		)
	}

	return recipes
}
