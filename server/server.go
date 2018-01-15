package server

import (
	"fmt"
	"net/http"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/db"
	"github.com/chvck/meal-planner/store"
)

func Run(cfg *config.Info) error {
	database := &db.SqlxAdapter{}
	err := database.Initialize(cfg.DbType, cfg.DbString)

	if err != nil {
		return err
	}

	store.StoreDb(database)

	address := fmt.Sprintf("%v:%v", cfg.Hostname, cfg.HttpPort)
	r := routes()

	return http.ListenAndServe(address, r)
}
