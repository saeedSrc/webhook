package app

import (
	"github.com/julienschmidt/httprouter"
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
	a.Logger.Info("Listening on port ", a.Config.WebHook.Port, "...")
	c := controller.NewController(a.Config, a.Logger)
	router := httprouter.New()
	router.POST("/api/v1/update", c.Update)
	err := http.ListenAndServe(":"+a.Config.WebHook.Port, router)
	if err != nil {
		panic(err)
	}
}
