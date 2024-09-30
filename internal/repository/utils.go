package repository

import (
	"database/sql"

	"github.com/AwesomeXjs/music-lib/internal/helpers"
)

// handleRollback - rollback transaction
func (s *SongRepo) handleRollback(tx *sql.Tx) error {
	if err := tx.Rollback(); err != nil {
		s.logger.Info(helpers.PgPrefix, helpers.FailedToRollback)
		return err
	}
	return nil
}

// executeInTransaction - execute function in transaction
func (s *SongRepo) executeInTransaction(execFn func(tx *sql.Tx) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	err = execFn(tx)
	if err != nil {
		err = s.handleRollback(tx)
		return err
	}

	return tx.Commit()
}
