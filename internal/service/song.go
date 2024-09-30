package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/internal/repository"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/google/uuid"
)

// SongService - interface for song service
type SongService struct {
	repo   *repository.Repository
	logger logger.Logger
	client *helpers.CustomClient
}

// NewSongService - create new song service
func NewSongService(repo *repository.Repository, logger logger.Logger, client *helpers.CustomClient) *SongService {
	return &SongService{
		repo:   repo,
		logger: logger,
		client: client,
	}
}

// CreateSong - create new song
func (s *SongService) CreateSong(input model.SongCreate) (string, error) {
	return s.repo.CreateSong(model.Song{
		ID:    uuid.New().String(),
		Group: input.Group,
		Song:  input.Song,
	})
}

// UpdateSong - update song
func (s *SongService) UpdateSong(id string, input model.SongUpdate) error {
	return s.repo.UpdateSong(id, input)
}

// DeleteSong - delete song
func (s *SongService) DeleteSong(id string) error {
	return s.repo.Song.DeleteSong(id)
}

// GetSongs - get songs
func (s *SongService) GetSongs(group, song, createdAt, text, link string, page, limit int) ([]model.Song, error) {
	offset := (page - 1) * limit
	return s.repo.Song.GetSongs(group, song, createdAt, text, link, offset, limit)
}

// GetVerse - get verse of the song
func (s *SongService) GetVerse(id string) (string, error) {
	return s.repo.Song.GetVerse(id)
}

// FetchSongData - fetch song data from mock service
func (s *SongService) FetchSongData(id string, input model.SongCreate) error {
	req, err := s.client.GetWithQuery(
		"/info",
		helpers.QueryParam{Key: "group", Value: input.Group},
		helpers.QueryParam{Key: "song", Value: input.Song})
	if err != nil {
		s.logger.Debug(helpers.PgPrefix, helpers.RequestError)
		return fmt.Errorf("failed to get song data: %v", err)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			s.logger.Debug(helpers.PgPrefix, err.Error())
			return
		}
	}(req.Body)

	var arr1 []model.SongUpdate
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		s.logger.Debug(helpers.PgPrefix, helpers.ReadBodyError)
		return fmt.Errorf("failed to read response body: %v", err)
	}

	err = json.Unmarshal(reqBody, &arr1)
	if err != nil {
		s.logger.Debug(helpers.PgPrefix, helpers.UnmarshalError)
		return fmt.Errorf("failed to unmarshal response body: %v", err)
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

// GetAllFromMockService - get all songs from mock service
func (s *SongService) GetAllFromMockService() ([]helpers.MockSongs, error) {
	req, err := s.client.Client.Get(s.client.SideServiceURL + "/info")
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			s.logger.Debug(helpers.PgPrefix, err.Error())
			return
		}
	}(req.Body)
	var data []helpers.MockSongs
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		s.logger.Debug(helpers.PgPrefix, helpers.ReadBodyError)
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	err = json.Unmarshal(reqBody, &data)

	if err != nil {
		s.logger.Debug(helpers.PgPrefix, helpers.UnmarshalError)
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return data, nil
}
