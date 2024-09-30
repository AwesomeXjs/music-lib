package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AwesomeXjs/music-lib/internal/db"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// SongRepo - repository for songs
type SongRepo struct {
	db     *sqlx.DB
	logger logger.Logger
}

// NewSongRepo - create new song repository
func NewSongRepo(db *sqlx.DB, logger logger.Logger) *SongRepo {
	return &SongRepo{
		db:     db,
		logger: logger,
	}
}

// CreateSong - create new song
func (s *SongRepo) CreateSong(song model.Song) (string, error) {

	if song.Text == "" && song.Link == "" && song.ReleaseDate == "" {
		song.Text = helpers.DefaultValueForFields
		song.Link = helpers.DefaultValueForFields
		song.ReleaseDate = helpers.DefaultValueForFields
	}

	var songID string
	err := s.executeInTransaction(func(tx *sql.Tx) error {
		query := fmt.Sprintf("INSERT INTO %s (id, group_name, song, text, link, release_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", db.SongsTable)
		result := tx.QueryRow(query, song.ID, song.Group, song.Song, song.Text, song.Link, song.ReleaseDate)
		err := result.Scan(&songID)
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, err.Error())
			return err
		}
		return err
	})

	if songID == "" || err != nil {
		return "", errors.New("Failed to create song or song already exists")
	}

	return songID, nil
}

// UpdateSong - update song
func (s *SongRepo) UpdateSong(id string, song model.SongUpdate) error {
	return s.executeInTransaction(func(tx *sql.Tx) error {
		setValues := make([]string, 0)
		args := make([]interface{}, 0)
		argID := 1
		for k, v := range map[string]*string{
			"song":         song.Song,
			"text":         song.Text,
			"link":         song.Link,
			"release_date": song.ReleaseDate,
			"group_name":   song.Group,
		} {
			if v != nil && len(*v) > 0 {
				setValues = append(setValues, fmt.Sprintf("%s=$%d", k, argID))
				args = append(args, *v)
				argID++
			}
		}

		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", db.SongsTable, setQuery, argID)

		stmt, err := tx.Prepare(query)
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Prepare failed")
			return fmt.Errorf("failed to prepare query: %w", err)
		}
		defer func(stmt *sql.Stmt) {
			if err = stmt.Close(); err != nil {
				s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Close failed")
				return
			}
		}(stmt)

		args = append(args, id)

		rows, err := stmt.Exec(args...)
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Exec failed")
			return fmt.Errorf("failed to update song: %w", err)
		}

		affected, err := rows.RowsAffected()
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Rows affected")
			return fmt.Errorf("failed to get rows affected: %w", err)
		}
		if affected == 0 {
			s.logger.Debug(helpers.PgPrefix, helpers.NoRowsAffected)
			return errors.New("no rows affected")
		}

		return nil
	})
}

// DeleteSong - delete song
func (s *SongRepo) DeleteSong(id string) error {
	return s.executeInTransaction(func(tx *sql.Tx) error {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", db.SongsTable)

		stmt, err := tx.Prepare(query)
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Prepare failed")
			return fmt.Errorf("failed to prepare query: %w", err)
		}
		defer func(stmt *sql.Stmt) {
			if err = stmt.Close(); err != nil {
				s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed+"Close failed")
				return
			}
		}(stmt)

		_, err = stmt.Exec(id)
		if err != nil {
			s.logger.Debug(helpers.PgPrefix, helpers.PgTransactionFailed)
			return fmt.Errorf("failed to delete song: %w", err)
		}
		return nil
	})
}

// GetSongs - get songs
func (s *SongRepo) GetSongs(group, song, createdAt, text, link string, offset, limit int) ([]model.Song, error) {
	query := fmt.Sprintf("SELECT id, group_name, song, text, link, release_date FROM %s", db.SongsTable)
	var args []interface{}
	argID := 1

	var conditions []string
	for k, v := range map[string]string{
		"group_name":   group,
		"song":         song,
		"text":         text,
		"link":         link,
		"release_date": createdAt,
	} {
		if v != "" {
			conditions = append(conditions, fmt.Sprintf("%s = $%d", k, argID))
			args = append(args, v)
			argID++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, limit, offset)

	rows, err := s.db.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			s.logger.Info(helpers.PgPrefix, helpers.FailedToClose)
			return
		}
	}(rows)

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		err = rows.StructScan(&song)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		songs = append(songs, song)
	}
	if songs == nil {
		return nil, errors.New("Songs not found")
	}
	return songs, nil
}

// GetVerse - get verse
func (s *SongRepo) GetVerse(id string) (string, error) {
	query := fmt.Sprintf("SELECT text FROM %s WHERE id = $1", db.SongsTable)

	var text string
	res, err := s.db.Query(query, id)

	if err != nil {
		s.logger.Info(helpers.PgPrefix, "Failed to get text from database, query error")
		return "", fmt.Errorf("query error: %v", err)
	}

	for res.Next() {
		err = res.Scan(&text)

		if err != nil {
			s.logger.Info(helpers.PgPrefix, "Failed to get text from database, scan error")
			return "", fmt.Errorf("scan error: %v", err)
		}
	}

	if res == nil {
		s.logger.Info(helpers.PgPrefix, "Failed to get text from database, result is nil")
		return "", fmt.Errorf("result is nil")
	}
	return text, nil
}
