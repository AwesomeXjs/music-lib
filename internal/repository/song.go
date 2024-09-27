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
	if song.Text == "" && song.Link == "" && song.ReleaseDate == "" {
		song.Text = helpers.DEFAULT_VALUE_FOR_FIELDS
		song.Link = helpers.DEFAULT_VALUE_FOR_FIELDS
		song.ReleaseDate = helpers.DEFAULT_VALUE_FOR_FIELDS
	}
	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, group_name, song, text, link, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", db.SONGS_TABLE)
	result := tx.QueryRow(query, song.Id, song.Group, song.Song, song.Text, song.Link, song.ReleaseDate)

	var songId string

	err = result.Scan(&songId)
	if songId == "" {
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(helpers.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return "", err
		}
		return "", errors.New(helpers.SONG_ALREADY_EXIST)
	}
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(helpers.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return "", err
		}
		return "", err
	}
	return songId, tx.Commit()
}

func (s *SongRepo) UpdateSong(id string, song model.SongUpdate) error {
	tx, err := s.db.Begin()
	if err != nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED)
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	for k, v := range map[string]*string{
		"song":         song.Song,
		"text":         song.Text,
		"link":         song.Link,
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
	rows, err := tx.Exec(query, args...)
	if rows == nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED+"Rows affected")
		return err
	}
	affected, err := rows.RowsAffected()
	if err != nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED+"Rows affected")
		return err
	}
	if affected == 0 {
		err = tx.Rollback()
		if err != nil {
			s.logger.Debug(helpers.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return err
		}
		return errors.New(helpers.CANNOT_FIND_ELEMENT_BY_ID + " or " + helpers.SONG_ALREADY_EXIST)
	}
	if err = tx.Commit(); err != nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_COMMIT_FAILED)
		return err
	}

	return nil
}

func (s *SongRepo) DeleteSong(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED)
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", db.SONGS_TABLE)
	_, err = tx.Exec(query, id)

	if err != nil {
		s.logger.Info(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED)
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(helpers.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return err
		}
	}
	return tx.Commit()
}

func (s *SongRepo) GetSongs(group, song, createdAt, text, link string, offset, limit int) ([]model.Song, error) {
	query := fmt.Sprintf("SELECT id, group_name, song, text, link, release_date FROM %s", db.SONGS_TABLE)
	var args []interface{}
	argId := 1

	// Добавляем условия фильтрации, если они есть
	var conditions []string

	for k, v := range map[string]string{
		"group_name":   group,
		"song":         song,
		"text":         text,
		"link":         link,
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
			s.logger.Info(helpers.PG_PREFIX, helpers.FAILED_TO_CLOSE)
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

func (s *SongRepo) GetVerse(id string) (string, error) {
	query := fmt.Sprintf("SELECT text FROM %s WHERE id = $1", db.SONGS_TABLE)

	var text string
	res, err := s.db.Query(query, id)

	for res.Next() {
		err = res.Scan(&text)

		if err != nil {
			s.logger.Info(helpers.PG_PREFIX, "Failed to get text from database, scan error")
			return "", err
		}
	}

	if res == nil {
		s.logger.Info(helpers.PG_PREFIX, "Failed to get text from database, result is nil")
		return "", err
	}
	return text, nil
}
