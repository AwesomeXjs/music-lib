package db

import (
	"fmt"
	"time"

	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Needed for migrations
	"github.com/jmoiron/sqlx"
)

const (
	// SongsTable - table songs name
	SongsTable = "songs"

	// KeepAlivePollPeriod - period for pinging db
	KeepAlivePollPeriod = 3 * time.Second

	// MaxTries - max tries to connect to db
	MaxTries = 20
)

// New - inits postgres db
func New(cfg *configs.Config, myLogger logger.Logger) (*sqlx.DB, error) {
	// get db url
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	// Open db connection
	db, err := sqlx.Open("postgres", databaseURL)
	fmt.Println(databaseURL)
	if err != nil {
		return nil, fmt.Errorf(" %v", err)
	}

	// First ping
	err = db.Ping()
	if err != nil {
		myLogger.Fatal(helpers.PgPrefix, "Failed first ping")
		return nil, fmt.Errorf(" %v", err)
	}
	myLogger.Info(helpers.PgPrefix, helpers.PgConnectSuccess)
	return db, nil
}

// KeepAlivePostgres keeps db alive
// if db disconnects, it will try to reconnect
func KeepAlivePostgres(database *sqlx.DB, myLogger logger.Logger) {
	count := 0
	for {
		time.Sleep(KeepAlivePollPeriod)
		err := database.Ping()
		if err != nil {
			count++
			if count == MaxTries {
				myLogger.Fatal(helpers.PgPrefix, helpers.DisconnectDB)
			}
			myLogger.Info(helpers.PgPrefix, helpers.ReconnectDB)
		}
	}
}

// MigrationUp - migrates db on start service
func MigrationUp(config *configs.Config, myLogger logger.Logger) error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
	m, err := migrate.New(
		"file://internal/db/migrations",
		dbURL)
	if m == nil || err != nil {
		myLogger.Fatal(helpers.PgPrefix, helpers.PgMigrateFailed)
		return err
	}
	err = m.Up()
	if err != nil {
		myLogger.Fatal(helpers.PgPrefix, helpers.PgMigrateFailed)
		return fmt.Errorf(" %v", err)
	}

	return nil
}
