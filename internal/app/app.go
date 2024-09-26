package app

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/internal/controller"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/internal/service"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type App struct {
	controller *controller.Controller
	service    *service.Service
	repository *repository.Repository

	Server *echo.Echo
}

func New(database interface{}, myLogger logger.Logger) *App {
	app := &App{}

	app.repository = repository.New(database, myLogger)
	app.service = service.New(app.repository, myLogger)
	app.controller = controller.New(app.service, myLogger)

	app.Server = echo.New()

	app.Server.Use(middleware.Recover())

	// handlers
	app.controller.InitRoutes(app.Server)

	return app
}

func (app *App) Run(myLogger logger.Logger, database interface{}) error {
	go func(myLogger logger.Logger) {
		fmt.Println("Server running...")
		err := app.Server.Start(os.Getenv("APP_PORT"))
		if err != nil {
			myLogger.Info(logger.APP_PREFIX, err.Error())
		}
	}(myLogger)

	err := app.gracefulShutdown(myLogger, database)
	if err != nil {
		return err
	}

	return nil
}