package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (e *Controller) addSong(ctx echo.Context) error {
	err := e.ResponseHelper(ctx, http.StatusOK, "Hello World", "", "", e.logger)
	if err != nil {
		return ctx.JSON(200, "JSON parse error")
	}

	ctx.JSON(200, "Hello World")
	return nil
}
