package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func (app *App) gracefulShutdown(myLogger logger.Logger, database *sqlx.DB) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit

	myLogger.Debug(helpers.AppPrefix, "Shutting down server..."+sig.String())

	if err := app.Server.Shutdown(context.Background()); err != nil {
		return err
	}
	if err := database.Close(); err != nil {
		myLogger.Debug(helpers.PgPrefix, helpers.DisconnectDB+" FAILED TO CLOSE DB")
		return err
	}
	return nil
}
