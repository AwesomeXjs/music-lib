package repository

import (
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type SongRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewSongRepo(db *sqlx.DB, logger logger.Logger) *SongRepo {
	return &SongRepo{
		db:     db,
		logger: logger,
	}
}
