package service

import (
	"encoding/json"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
)

type SongService struct {
	repo           *repository.Repository
	logger         logger.Logger
	sideServiceUrl string
	client         CustomClient
}

func NewSongService(repo *repository.Repository, logger logger.Logger, sideServiceUrl string, client CustomClient) *SongService {
	return &SongService{
		repo:           repo,
		logger:         logger,
		sideServiceUrl: sideServiceUrl,
		client:         client,
	}
}

func (s *SongService) CreateSong(input model.SongCreate) error {
	req, err := s.client.GetWithQuery(s.sideServiceUrl,
		"/info",
		QueryParam{Key: "group", Value: input.Group},
		QueryParam{Key: "song", Value: input.Song})
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, err.Error())
		}
	}(req.Body)

	var arr1 []model.SongRequest
	reqBody, err := io.ReadAll(req.Body)
	err = json.Unmarshal(reqBody, &arr1)
	if len(arr1) == 0 {
		return errors.New(helpers.SONGS_NOT_FOUND)
	}

	return s.repo.CreateSong(model.Song{
		Id:          uuid.New().String(),
		Group:       input.Group,
		Song:        input.Song,
		Text:        arr1[0].Text,
		Patronymic:  arr1[0].Patronymic,
		ReleaseDate: arr1[0].ReleaseDate,
	})
}

func (s *SongService) UpdateSong(id string, input model.SongUpdate) error {
	return s.repo.UpdateSong(id, input)
}

func (s *SongService) DeleteSong(id string) error {
	return s.repo.Song.DeleteSong(id)
}
