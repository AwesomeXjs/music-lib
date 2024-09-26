package repository

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/internal/db"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
)

type SongRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

func NewSongRepo(db *sqlx.DB, logger logger.Logger) *SongRepo {
	return &SongRepo{
		db:     db,
		logger: logger,
	}
}

func (s *SongRepo) CreateSong(song model.Song) (string, error) {
	if song.Text == "" && song.Patronymic == "" && song.ReleaseDate == "" {
		song.Text = "NOT FOUND"
		song.Patronymic = "NOT FOUND"
		song.ReleaseDate = "NOT FOUND"
	}

	fmt.Printf("%+v", song)
	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, group_name, song, text, patronymic, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", db.SONGS_TABLE)
	result := tx.QueryRow(query, song.Id, song.Group, song.Song, song.Text, song.Patronymic, song.ReleaseDate)

	var songId string

	err = result.Scan(&songId)
	if songId == "" {
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return "", err
		}
		return "", errors.New("Song already exists")
	}
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return "", err
		}
		return "", err
	}
	fmt.Println("SONG ID", songId)
	return songId, tx.Commit()
}

func (s *SongRepo) UpdateSong(id string, song model.SongUpdate) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	for k, v := range map[string]*string{
		"song":         song.Song,
		"text":         song.Text,
		"patronymic":   song.Patronymic,
		"release_date": song.ReleaseDate,
		"group_name":   song.Group,
	} {
		if v != nil && len(*v) > 0 {
			setValues = append(setValues, fmt.Sprintf("%s=$%d", k, argId))
			args = append(args, *v)
			argId++
		}
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", db.SONGS_TABLE, setQuery, argId)
	args = append(args, id)
	_, err = tx.Exec(query, args...)

	if err != nil {
		s.logger.Info(logger.PG_PREFIX, err.Error())
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return err
		}
	}
	return tx.Commit()
}

func (s *SongRepo) DeleteSong(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", db.SONGS_TABLE)
	_, err = tx.Exec(query, id)

	if err != nil {
		s.logger.Info(logger.PG_PREFIX, err.Error())
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return err
		}
	}
	return tx.Commit()
}

func (s *SongRepo) GetSongs(group, song, createdAt, text, patronymic string, offset, limit int) ([]model.Song, error) {
	query := fmt.Sprintf("SELECT id, group_name, song, text, patronymic, release_date FROM %s", db.SONGS_TABLE)
	var args []interface{}
	argId := 1

	// Добавляем условия фильтрации, если они присутствуют
	var conditions []string

	for k, v := range map[string]string{
		"group_name":   group,
		"song":         song,
		"text":         text,
		"patronymic":   patronymic,
		"release_date": createdAt,
	} {
		if v != "" {
			conditions = append(conditions, fmt.Sprintf("%s = $%d", k, argId))
			args = append(args, v)
			argId++
		}
	}

	// Если есть условия, добавляем их в запрос
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Добавляем LIMIT и OFFSET
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, limit, offset)

	// Выполняем запрос
	rows, err := s.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_CLOSE)
		}
	}(rows)

	// Обрабатываем результаты
	var songs []model.Song
	for rows.Next() {
		var song model.Song
		err = rows.StructScan(&song)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	if songs == nil {
		return nil, errors.New("Songs not found")
	}
	return songs, nil
}
