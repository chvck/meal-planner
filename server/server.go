package server

import (
	"fmt"
	"net/http"

	"github.com/chvck/meal-planner/controller"
	"github.com/chvck/meal-planner/datamodel"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/service"
)

// Run is the entry point for running the server
func Run(cfg *config.Info) error {
	database := &db.SqlxAdapter{}
	err := database.Initialize(cfg.DbType, cfg.DbString)
	if err != nil {
		return err
	}

	userDataModel := datamodel.NewSQLUser(database)
	recipeDataModel := datamodel.NewSQLRecipe(database)
	menuDataModel := datamodel.NewSQLMenu(database)
	// plannerDataModel := datamodel.NewPlannerRecipe(database)

	userService := service.NewUserService(userDataModel)
	recipeService := service.NewRecipeService(recipeDataModel)
	menuService := service.NewMenuService(menuDataModel, recipeDataModel)
	// plannerService := service.NewPlannerService(plannerDataModel, menuDataModel, recipeDataModel)

	menuController := controller.NewMenuController(menuService)
	userController := controller.NewUserController(userService)
	recipeController := controller.NewRecipeController(recipeService)

	handler := NewHandler(menuController, recipeController, userController)

	r := routes(handler)

	address := fmt.Sprintf("%v:%v", cfg.Hostname, cfg.HTTPPort)
	return http.ListenAndServe(address, r)
}
