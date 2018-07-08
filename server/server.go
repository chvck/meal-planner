package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chvck/meal-planner/config"
	"github.com/chvck/meal-planner/controller"
	"github.com/chvck/meal-planner/datastore/cbdatastore"
)

// Run is the entry point for running the server
func Run(cfg *config.Info) (*http.Server, error) {
	dataStore, err := cbdatastore.NewCBDataStore(
		cfg.DbServer,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbUsername,
		cfg.DbPassword,
	)
	if err != nil {
		return nil, err
	}

	cont := controller.NewStandardController(dataStore, cfg.AuthKey)

	handler := NewHandler(cont)

	r := routes(handler, cfg.AuthKey)

	address := fmt.Sprintf("%v:%v", cfg.Hostname, cfg.HTTPPort)

	srv := &http.Server{Addr: address, Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Error running server: %s", err)
			return
		}
	}()

	return srv, nil
}
