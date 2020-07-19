package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"Sparkle/config"
)

type DB struct {
	Pool *sql.DB
	Tx   *sql.Tx
}

type Performer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func (db DB) Performer() Performer {
	if db.Tx != nil {
		return db.Tx
	}
	return db.Pool
}

func Init(config config.DatabaseConfig) (DB, error) {

	var db DB

	pool, err := sql.Open("postgres", config.Url)

	if err != nil {
		return db, err
	}

	pool.SetConnMaxLifetime(config.ConnLifetime)
	pool.SetMaxOpenConns(config.OpenConns)
	pool.SetMaxIdleConns(config.IdleConns)

	if err := pool.Ping(); err != nil {
		return db, err
	}

	db.Pool = pool

	return db, nil

}
