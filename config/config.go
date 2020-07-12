package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Port         string        `envconfig:"port"`
	ReadTimeout  time.Duration `envconfig:"read_timeout"`
	WriteTimeout time.Duration `envconfig:"write_timeout"`
	IdleTimeout  time.Duration `envconfig:"idle_timeout"`
}

type DatabaseConfig struct {
	Redis    RedisConfig    `envconfig:"redis"`
	Postgres PostgresConfig `envconfig:"postgres"`
}

type RedisConfig struct {
	Addr     string `envconfig:"addr"`
	Password string `envconfig:"password"`
	DB       int    `envconfig:"db"`
}

type PostgresConfig struct {
	Url          string        `envconfig:"url"`
	OpenConns    int           `envconfig:"open_conns"`
	IdleConns    int           `envconfig:"idle_conns"`
	ConnLifetime time.Duration `envconfig:"conn_lifetime"`
}

type EnvironmentConfig struct {
	Server   ServerConfig   `envconfig:"server"`
	Database DatabaseConfig `envconfig:"db"`
}

func LoadConfig() (EnvironmentConfig, error) {

	var config EnvironmentConfig

	err := envconfig.Process("sparkle", &config)

	if err != nil {
		return config, err
	}

	return config, nil

}
