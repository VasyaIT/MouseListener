package main

import (
	"MouseListener/config"
	"MouseListener/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig("configuration.yaml")
	if err != nil {
		log.Fatal("Config error: ", err)
	}
	app.Run(cfg)
}
