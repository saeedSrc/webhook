package main

import (
	"webhook/app"
	"webhook/config"
)

func main() {
	cfg := config.Init("./config.yaml")
	application := app.NewApp()
	application.Init(cfg)
}
