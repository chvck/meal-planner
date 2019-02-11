package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run is the entry point for running the server
func Run(cfg *Info) (*http.Server, error) {
	dataStore, err := NewDataStore(
		cfg.DbServer,
		cfg.DbPort,
		cfg.DbName,
		cfg.DbUsername,
		cfg.DbPassword,
	)
	if err != nil {
		return nil, err
	}

	cont := NewStandardController(dataStore, cfg.AuthKey)

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

