package database

import (
	"database/sql"

	"context"

	_ "github.com/lib/pq"

	redis "github.com/go-redis/redis/v8"

	"Sparkle/config"
)

type DB struct {
	pool  *sql.DB
	cache *redis.Client
	tx    *sql.Tx
	ctx   context.Context
}

type Performer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func (db DB) performer() Performer {
	if db.tx != nil {
		return db.tx
	}
	return db.pool
}

func Init(config config.DatabaseConfig) (DB, error) {

	var db DB

	db.ctx = context.Background()

	pool, err := sql.Open("postgres", config.Postgres.Url)

	if err != nil {
		return db, err
	}

	pool.SetConnMaxLifetime(config.Postgres.ConnLifetime)
	pool.SetMaxOpenConns(config.Postgres.OpenConns)
	pool.SetMaxIdleConns(config.Postgres.IdleConns)

	if err := pool.Ping(); err != nil {
		return db, err
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	if _, err := cache.Ping(db.ctx).Result(); err != nil {
		return db, err
	}

	db.pool = pool
	db.cache = cache

	return db, nil

}
