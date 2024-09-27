package helpers

import (
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type Response struct {
	Title   string `json:"title"`
	Detail  string `json:"detail"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

func ResponseHelper(ctx echo.Context, statusCode int, message, detail, request string, myLogger logger.Logger) error {
	myLogger.Response(RESPONSE_PREFIX, strconv.Itoa(statusCode), message+" "+detail)
	err := ctx.JSON(statusCode, Response{Title: message, Detail: detail, Request: request, Time: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		myLogger.Info(RESPONSE_PREFIX, err.Error())
		return err
	}
	return nil
}
