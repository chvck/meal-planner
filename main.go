package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	conf, err := Load("config.json")

	if err != nil {
		log.Fatalln(err)
	}

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	srv, err := Run(conf)
	if err != nil {
		log.Fatal(err)
	}

	<-stop
	log.Println("Stopping server")
	srv.Shutdown(nil)
}

