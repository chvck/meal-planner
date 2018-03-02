package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/controller"
	"github.com/chvck/meal-planner/datamodel"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/service"
)

// Run is the entry point for running the server
func Run(cfg *config.Info) (*http.Server, error) {
	database := &db.SqlxAdapter{}
	err := database.Initialize(cfg.DbType, cfg.DbString)
	if err != nil {
		return nil, err
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

	srv := &http.Server{Addr: address, Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Error running server: %s", err)
		}
	}()

	return srv, nil
}
