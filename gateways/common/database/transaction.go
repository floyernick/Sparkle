package database

import (
	"Sparkle/tools/logger"
)

func (db *DB) Transaction() error {
	if db.Tx != nil {
		return nil
	}

	tx, err := db.Pool.Begin()

	if err != nil {
		logger.Warning(err)
		return err
	}

	db.Tx = tx

	return nil

}

func (db DB) Commit() error {

	if db.Tx == nil {
		return nil
	}

	err := db.Tx.Commit()

	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil

}
