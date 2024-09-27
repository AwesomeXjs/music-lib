package db

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	SONGS_TABLE         = "songs"
	KeepAlivePollPeriod = 3 * time.Second
	MaxTries            = 1200 // 1 hour
)

func New(cfg *configs.Config, myLogger logger.Logger) (*sqlx.DB, error) {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		myLogger.Fatal(helpers.PG_PREFIX, "Failed first ping")
		return nil, err
	}
	myLogger.Info(helpers.PG_PREFIX, helpers.PG_CONNECT_SUCCESS)
	return db, nil
}

func KeepAlivePostgres(database *sqlx.DB, myLogger logger.Logger) {
	count := 0
	for {
		time.Sleep(KeepAlivePollPeriod)
		err := database.Ping()
		if err != nil {
			count++
			if count == MaxTries {
				myLogger.Fatal(helpers.PG_PREFIX, helpers.DISCONNECT_DB)
			}
			myLogger.Info(helpers.PG_PREFIX, helpers.RECONECT_DB)
		}
	}
}
