package main

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/internal/app"
	"github.com/AwesomeXjs/music-lib/internal/db"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	zaplogger "github.com/AwesomeXjs/music-lib/pkg/logger/zap"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title Music library API
// @version 1.0
// @description API Server for Music library application
// @host localhost:8080
// @BasePath /
// @in header
func main() {

	// logger init
	myLogger := zaplogger.New()

	// init env config
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		myLogger.Fatal("[ ENV ]", "failed to load env variables")
	}

	config := configs.New()

	postgres, err := db.New(config, myLogger)
	if err != nil {
		myLogger.Fatal(logger.PG_PREFIX, logger.PG_CONNECTION_FAILED)
	}

	// Keep Alive Postgres
	go db.KeepAlivePostgres(postgres, myLogger)

	// init new app
	myApp := app.New(postgres, myLogger, config)

	// start server
	err = myApp.Run(myLogger, postgres)
	if err != nil {
		myLogger.Fatal(logger.APP_PREFIX, err.Error())
	}
}
