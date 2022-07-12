package config

import (
	"context"
	"github.com/go-redis/redis"
)

func NewConnectRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client

}
