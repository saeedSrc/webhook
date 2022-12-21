package app

import (
	"go.uber.org/zap"
	"log"
	"net/http"
	"webhook/config"
	"webhook/controller"
)

type App struct {
	Logger *zap.SugaredLogger
	Config *config.Config
}

func NewApp() *App {
	app := &App{}
	return app
}

func (a *App) Init(cfg *config.Config) {
	a.Config = cfg
	a.Logger = a.initialLogger()
	a.registerRoutes()
}

func (a *App) initialLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	return logger.Sugar()
}

func (a *App) registerRoutes() {
	c := controller.NewController(a.Config, a.Logger)
	http.HandleFunc("/api/v1/update", c.Update)
	a.Logger.Info("Listening on port", a.Config.WebHook.Port, ".")
	if err := http.ListenAndServe(":"+a.Config.WebHook.Port, nil); err != nil {
		panic(err)
	}
}
