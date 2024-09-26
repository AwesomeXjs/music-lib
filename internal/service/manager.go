package service

import (
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
)

type Song interface {
	CreateSong(input model.SongCreate) (string, error)
	UpdateSong(id string, input model.SongUpdate) error
	DeleteSong(id string) error
	GetSongs(group, song, createdAt, text, patronymic string, page, limit int) ([]model.Song, error)
}

type Service struct {
	Song
}

func New(repo *repository.Repository, logger logger.Logger, mockUrl string) *Service {
	client := CustomClient{client: NewClient()}
	return &Service{
		Song: NewSongService(repo, logger, mockUrl, client),
	}
}
