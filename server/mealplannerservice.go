package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"

	"github.com/chvck/meal-planner/proto/model"
	"github.com/chvck/meal-planner/proto/service"
)

type MealPlannerService struct {
	datastore DataStore
	authKey   string
}

func (mps *MealPlannerService) AllRecipes(ctx context.Context, query *service.AllRecipesRequest) (*service.AllRecipesResponse, error) {
	var recipes []*model.Recipe
	met, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(met)
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
	u := mps.datastore.UserValidatePassword(query.Username, []byte(query.Password))
	if u == nil {
		return nil, errors.New("invalid user credentials provided")
	}

	t, err := mps.createToken(u)
	if err != nil {
		return nil, err
	}

	return &service.LoginUserResponse{
		Token: t,
	}, nil
}

func (mps *MealPlannerService) createToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"id":        user.Id,
		"lastLogin": user.LastLogin,
		"level":     model.LevelUser,
	})
	return token.SignedString([]byte(mps.authKey))
}

func (mps *MealPlannerService) validateUser(token string, reqLevel float64, u *model.User) (*model.User, error) {
	bearerToken := strings.Split(token, " ")
	if len(bearerToken) == 2 {
		token, err := jwt.Parse(bearerToken[1], mps.parseToken)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["level"].(float64) < reqLevel {
				return nil, errors.Wrap(err, "Not Authorized")
			}
			u := &model.User{Id: claims["id"].(string), Username: claims["username"].(string)}
			return u, nil
		}

		return nil, errors.Wrap(err, "Invalid token")
	}

	return nil, errors.New("Invalid token")
}

func (mps *MealPlannerService) parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("there was an error")
	}
	return []byte(mps.authKey), nil
}
