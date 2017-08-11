package main

import (
	"log"

	"github.com/bunk/vitable/ingress/api"
	"github.com/bunk/vitable/ingress/config"
)

func main() {
	config := config.Get()
	log.Printf("Running in %s mode", config.Mode)

	err := api.Start(config)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
