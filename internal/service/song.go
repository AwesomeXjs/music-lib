package service

import (
	"encoding/json"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/google/uuid"
	"io"
)

type SongService struct {
	repo           *repository.Repository
	logger         logger.Logger
	sideServiceUrl string
	client         *helpers.CustomClient
}

func NewSongService(repo *repository.Repository, logger logger.Logger, client *helpers.CustomClient) *SongService {
	return &SongService{
		repo:   repo,
		logger: logger,
		client: client,
	}
}

func (s *SongService) CreateSong(input model.SongCreate) (string, error) {
	return s.repo.CreateSong(model.Song{
		Id:    uuid.New().String(),
		Group: input.Group,
		Song:  input.Song,
	})
}

func (s *SongService) UpdateSong(id string, input model.SongUpdate) error {
	return s.repo.UpdateSong(id, input)
}

func (s *SongService) DeleteSong(id string) error {
	return s.repo.Song.DeleteSong(id)
}

func (s *SongService) GetSongs(group, song, createdAt, text, link string, page, limit int) ([]model.Song, error) {
	offset := (page - 1) * limit
	return s.repo.Song.GetSongs(group, song, createdAt, text, link, offset, limit)
}

func (s *SongService) GetVerse(id string) (string, error) {
	return s.repo.Song.GetVerse(id)
}

func (s *SongService) FetchSongData(id string, input model.SongCreate) error {
	req, err := s.client.GetWithQuery(
		"/info",
		helpers.QueryParam{Key: "group", Value: input.Group},
		helpers.QueryParam{Key: "song", Value: input.Song})
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var arr1 []model.SongUpdate
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(reqBody, &arr1)
	if err != nil {
		return err
	}
	// Если данные найдены, обновляем песню в базе
	if len(arr1) > 0 {
		return s.repo.Song.UpdateSong(id, model.SongUpdate{
			Text:        arr1[0].Text,
			Link:        arr1[0].Link,
			ReleaseDate: arr1[0].ReleaseDate,
		})
	}

	return nil
}

func (s *SongService) GetAllFromMockService() ([]helpers.MockSongs, error) {
	req, err := s.client.Client.Get(s.client.SideServiceUrl + "/info")
	defer req.Body.Close()
	var data []helpers.MockSongs
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
