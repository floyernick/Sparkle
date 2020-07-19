package cache

import (
	"context"

	redis "github.com/go-redis/redis/v8"

	"Sparkle/config"
)

type Cache struct {
	Client *redis.Client
	Ctx    context.Context
}

func Init(config config.CacheConfig) (Cache, error) {

	var cache Cache

	cache.Ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	if _, err := client.Ping(cache.Ctx).Result(); err != nil {
		return cache, err
	}

	cache.Client = client

	return cache, nil

}
