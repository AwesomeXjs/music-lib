package helpers

import (
	"strconv"
	"time"

	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Response struct
type Response struct {
	Title   string `json:"title"`
	Detail  string `json:"detail"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

// ResponseHelper helper function
func ResponseHelper(ctx echo.Context, statusCode int, message, detail string, myLogger logger.Logger) error {
	myLogger.Response(ResponsePrefix, strconv.Itoa(statusCode), message+" "+detail)
	err := ctx.JSON(statusCode, Response{Title: message, Detail: detail, Request: ctx.Request().RequestURI, Time: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		myLogger.Info(ResponsePrefix, err.Error())
		return err
	}
	return nil
}
