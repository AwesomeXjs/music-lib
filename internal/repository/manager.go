package repository

import (
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Song interface {
	CreateSong(input model.Song) error
	UpdateSong(id string, input model.SongUpdate) error
	DeleteSong(id string) error
}

type Repository struct {
	Song
}

func New(db interface{}, myLogger logger.Logger) *Repository {
	// проверка типа чтобы мы не зависели от одной базы данных и могли легко переключить репозиторий на MongoDB
	switch database := db.(type) {
	case *sqlx.DB:
		return &Repository{
			Song: NewSongRepo(database, myLogger),
		}

	default:
		myLogger.Fatal(logger.REPO_PREFIX, logger.REPO_CREATE_FAILED)
		return nil
	}
}
