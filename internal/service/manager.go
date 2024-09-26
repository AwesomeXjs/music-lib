package service

import (
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
)

type Song interface {
	CreateSong(input model.SongCreate) error
	UpdateSong(id string, input model.SongUpdate) error
	DeleteSong(id string) error
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
