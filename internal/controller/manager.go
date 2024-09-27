package controller

import (
	_ "github.com/AwesomeXjs/music-lib/docs"
	"github.com/AwesomeXjs/music-lib/internal/service"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

type Controller struct {
	service *service.Service
	logger  logger.Logger
}

func New(service *service.Service, logger logger.Logger) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

func (e *Controller) InitRoutes(server *echo.Echo) {
	// Swagger init
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	// App routes
	v1 := server.Group("/v1")
	{
		// song routes
		songs := v1.Group("/songs")
		{
			songs.POST("", e.CreateSong)
			songs.GET("", e.GetSongs)
			songs.PUT("/:id", e.UpdateSong)
			songs.DELETE("/:id", e.DeleteSong)
			songs.GET("/verse/:id", e.GetVerse)
		}
		v1.GET("/all", e.GetAll)
	}
}
