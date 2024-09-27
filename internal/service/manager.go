package service

import (
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
)

type Song interface {
	CreateSong(input model.SongCreate) (string, error)
	UpdateSong(id string, input model.SongUpdate) error
	DeleteSong(id string) error
	GetSongs(group, song, createdAt, text, link string, page, limit int) ([]model.Song, error)
	GetVerse(id string) (string, error)
	FetchSongData(id string, input model.SongCreate) error
	GetAllFromMockService() ([]helpers.MockSongs, error)
}

type Service struct {
	Song
}

func New(repo *repository.Repository, logger logger.Logger) *Service {
	client := helpers.NewCustomClient(logger)
	return &Service{
		Song: NewSongService(repo, logger, client),
	}
}
