package main

import (
	"github.com/chvck/meal-planner/config"
	"log"
	"github.com/chvck/meal-planner/server"
)

func main() {
	conf, err := config.Load("config.json")

	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(server.Run(conf))
}
