package repository

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/internal/db"
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/AwesomeXjs/music-lib/internal/model"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"github.com/jmoiron/sqlx"
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

func (s *SongRepo) CreateSong(song model.Song) error {
	fmt.Printf("%+v", song)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, group_name, song, text, patronymic, release_date) VALUES ($1, $2, $3, $4, $5, $6)", db.SONGS_TABLE)
	_, err = tx.Exec(query, song.Id, song.Group, song.Song, song.Text, song.Patronymic, song.ReleaseDate)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			s.logger.Info(logger.PG_PREFIX, helpers.FAILED_TO_ROLLBACK)
			return err
		}
		return err
	}
	return tx.Commit()
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
