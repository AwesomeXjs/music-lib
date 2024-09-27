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
	GetSongs(group, song, createdAt, text, link string, offset, limit int) ([]model.Song, error)
	GetVerse(id string) (string, error)
}

type Repository struct {
	Song
}

func New(db *sqlx.DB, myLogger logger.Logger) *Repository {
	return &Repository{
		Song: NewSongRepo(db, myLogger),
	}
}
