package repository

import (
	"database/sql"
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

	var songId string
	err := s.executeInTransaction(func(tx *sql.Tx) error {
		query := fmt.Sprintf("INSERT INTO %s (id, group_name, song, text, link, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", db.SONGS_TABLE)
		result := tx.QueryRow(query, song.Id, song.Group, song.Song, song.Text, song.Link, song.ReleaseDate)
		err := result.Scan(&songId)
		if err != nil {
			s.logger.Debug(helpers.PG_PREFIX, err.Error())
			return err
		}
		return err
	})

	if songId == "" {
		return "", errors.New("Failed to create song or song already exists")
	}
	if err != nil {
		return "", err
	}
	return songId, nil
}

func (s *SongRepo) UpdateSong(id string, song model.SongUpdate) error {
	return s.executeInTransaction(func(tx *sql.Tx) error {
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
		if err != nil {
			s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED+"Exec failed")
			return err
		}

		affected, err := rows.RowsAffected()
		if err != nil {
			s.logger.Debug(helpers.PG_PREFIX, helpers.PG_TRANSACTION_FAILED+"Rows affected")
			return err
		}
		if affected == 0 {
			s.logger.Debug(helpers.PG_PREFIX, helpers.NO_ROWS_AFFECTED)
			return errors.New("no rows affected")
		}

		return nil
	})
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
		return s.handleRollback(tx)
	}

	if err = tx.Commit(); err != nil {
		s.logger.Debug(helpers.PG_PREFIX, helpers.PG_COMMIT_FAILED)
		return err
	}
	return nil
}

func (s *SongRepo) GetSongs(group, song, createdAt, text, link string, offset, limit int) ([]model.Song, error) {
	query := fmt.Sprintf("SELECT id, group_name, song, text, link, release_date FROM %s", db.SONGS_TABLE)
	var args []interface{}
	argId := 1

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

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, limit, offset)

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
