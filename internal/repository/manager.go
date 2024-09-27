package repository

import (
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Song interface {
	CreateSong(input model.Song) (string, error)
	UpdateSong(id string, input model.SongUpdate) error
	DeleteSong(id string) error
	GetSongs(group, song, createdAt, text, patronymic string, offset, limit int) ([]model.Song, error)
	GetVerse(id string) (string, error)
}

type Repository struct {
	Song
}

func New(db interface{}, myLogger logger.Logger) *Repository {
	// проверка типа чтобы мы не зависели от одной базы данных и могли легко переключится на репозиторий с другой бд
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
