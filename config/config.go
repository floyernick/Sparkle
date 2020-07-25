package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Port           string        `envconfig:"port"`
	ReadTimeout    time.Duration `envconfig:"read_timeout"`
	WriteTimeout   time.Duration `envconfig:"write_timeout"`
	IdleTimeout    time.Duration `envconfig:"idle_timeout"`
	HandlerTimeout time.Duration `envconfig:"handler_timeout"`
}

type CacheConfig struct {
	Addr     string `envconfig:"addr"`
	Password string `envconfig:"password"`
	DB       int    `envconfig:"db"`
}

type DatabaseConfig struct {
	Url          string        `envconfig:"url"`
	OpenConns    int           `envconfig:"open_conns"`
	IdleConns    int           `envconfig:"idle_conns"`
	ConnLifetime time.Duration `envconfig:"conn_lifetime"`
}

type LoggerConfig struct {
	Debug bool `envconfig:"debug"`
}

type EnvironmentConfig struct {
	Server   ServerConfig   `envconfig:"server"`
	Database DatabaseConfig `envconfig:"db"`
	Cache    CacheConfig    `envconfig:"cache"`
	Logger   LoggerConfig   `envconfig:"logger"`
}

func LoadConfig() (EnvironmentConfig, error) {

	var config EnvironmentConfig

	err := envconfig.Process("sparkle", &config)

	if err != nil {
		return config, err
	}

	return config, nil

}
