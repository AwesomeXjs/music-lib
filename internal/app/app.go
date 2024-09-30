package app

import (
	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/internal/controller"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/internal/service"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App - main struct for app
type App struct {
	controller *controller.Controller
	service    *service.Service
	repository *repository.Repository

	Server *echo.Echo
	config *configs.Config
}

// New - create new app
func New(database *sqlx.DB, myLogger logger.Logger, config *configs.Config) *App {
	// Init app
	app := &App{}
	app.config = config
	app.repository = repository.New(database, myLogger)
	app.service = service.New(app.repository, myLogger)
	app.controller = controller.New(app.service, myLogger)
	app.Server = echo.New()

	// MW
	app.Server.Use(middleware.Recover())
	app.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:9999"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// ROUTES
	app.controller.InitRoutes(app.Server)

	return app
}

// Run - start service
func (app *App) Run(myLogger logger.Logger, database *sqlx.DB) error {
	go func(myLogger logger.Logger) {
		myLogger.Info("SERVER", "Server running...")
		err := app.Server.Start(app.config.AppPort)
		if err != nil {
			myLogger.Debug(helpers.AppPrefix, err.Error())
		}
	}(myLogger)

	err := app.gracefulShutdown(myLogger, database)
	if err != nil {
		myLogger.Debug(helpers.AppPrefix, err.Error())
		return err
	}
	return nil
}
