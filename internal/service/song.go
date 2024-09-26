package service

import (
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
)

type SongService struct {
	repo   *repository.Repository
	logger logger.Logger
}

func NewSongService(repo *repository.Repository, logger logger.Logger) *SongService {
	return &SongService{
		repo:   repo,
		logger: logger,
	}
}
