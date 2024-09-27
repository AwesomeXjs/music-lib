package main

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/internal/app"
	"github.com/AwesomeXjs/music-lib/internal/db"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	zaplogger "github.com/AwesomeXjs/music-lib/pkg/logger/zap"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title Music library API
// @version 1.0
// @description API Server for Music library application
// @host localhost:9999
// @BasePath /
// @in header
func main() {

	// logger init
	myLogger := zaplogger.New()

	// init env config
	if err := godotenv.Load(); err != nil {
		myLogger.Fatal("[ ENV ]", "failed to load env variables")
	}

	config := configs.New(myLogger)
	postgres, err := db.New(config, myLogger)
	if err != nil {
		myLogger.Fatal(helpers.PG_PREFIX, helpers.PG_CONNECTION_FAILED)
	}

	// Keep Alive Postgres
	go db.KeepAlivePostgres(postgres, myLogger)

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
	m, err := migrate.New(
		"file://internal/db/migrations",
		dbUrl)
	if err != nil {
		fmt.Println(err)
	}
	err = m.Up()
	if err != nil {
		myLogger.Fatal(helpers.PG_PREFIX, helpers.PG_MIGRATION_FAILED)
		return
	}

	// init new app
	myApp := app.New(postgres, myLogger, config)

	// start server
	err = myApp.Run(myLogger, postgres)
	if err != nil {
		myLogger.Fatal(helpers.APP_PREFIX, err.Error())
	}
}
