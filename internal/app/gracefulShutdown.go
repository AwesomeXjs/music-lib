package app

import (
	"context"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"syscall"
)

func (app *App) gracefulShutdown(myLogger logger.Logger, database interface{}) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit

	myLogger.Info(logger.APP_PREFIX, "Shutting down server..."+sig.String())

	if err := app.Server.Shutdown(context.Background()); err != nil {
		return err
	}

	switch v := database.(type) {
	case *sqlx.DB:
		if err := v.Close(); err != nil {
			return err
		}
	default:
		myLogger.Info(logger.PG_PREFIX, logger.DISCONNECT_DB+" FAILED TO CLOSE DB")
		return nil
	}
	return nil
}
