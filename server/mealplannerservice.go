package server

import (
	"context"
	"github.com/chvck/meal-planner/proto/model"
	"github.com/chvck/meal-planner/proto/service"
)

type MealPlannerService struct {
	datastore DataStore
}

func (mps *MealPlannerService) AllRecipes(ctx context.Context, query *service.AllRecipesRequest) (*service.AllRecipesResponse, error) {
	var recipes []*model.Recipe
	return &service.AllRecipesResponse{Recipes: recipes}, nil
}

func (mps *MealPlannerService) RecipeByID(ctx context.Context, query *service.RecipeByIDRequest) (*service.RecipeByIDResponse, error) {
	return &service.RecipeByIDResponse{}, nil
}

func (mps *MealPlannerService) CreateRecipe(ctx context.Context, query *service.CreateRecipeRequest) (*service.CreateRecipeResponse, error) {
	return &service.CreateRecipeResponse{}, nil
}

func (mps *MealPlannerService) UpdateRecipe(ctx context.Context, query *service.UpdateRecipeRequest) (*service.UpdateRecipeResponse, error) {
	return &service.UpdateRecipeResponse{}, nil
}

func (mps *MealPlannerService) DeleteRecipe(ctx context.Context, query *service.DeleteRecipeRequest) (*service.DeleteRecipeResponse, error) {
	return &service.DeleteRecipeResponse{}, nil
}

func (mps *MealPlannerService) CreateUser(ctx context.Context, query *service.CreateUserRequest) (*service.CreateUserResponse, error) {
	return &service.CreateUserResponse{}, nil
}

func (mps *MealPlannerService) LoginUser(ctx context.Context, query *service.LoginUserRequest) (*service.LoginUserResponse, error) {
	return &service.LoginUserResponse{}, nil
}
