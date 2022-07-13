package config

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewConnectRedis(ctx context.Context) *redis.Client {

	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	return client

}
