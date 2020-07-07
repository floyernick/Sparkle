package config

import (
	"os"
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
	Url          string        `envconfig:"url"`
	OpenConns    int           `envconfig:"open_conns"`
	IdleConns    int           `envconfig:"idle_conns"`
	ConnLifetime time.Duration `envconfig:"conn_lifetime"`
}

type EnvironmentConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func LoadConfig() (EnvironmentConfig, error) {

	var config EnvironmentConfig

	prefix := "sparkle_" + os.Getenv("SPARKLE_ENV")

	err := envconfig.Process(prefix+"_server", &config.Server)

	if err != nil {
		return config, err
	}

	err = envconfig.Process(prefix+"_database", &config.Database)

	if err != nil {
		return config, err
	}

	return config, nil

}
