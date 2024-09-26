package service

import (
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
)

type Song interface {
}

type Service struct {
	Song
}

func New(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		Song: NewSongService(repo, logger),
	}
}
