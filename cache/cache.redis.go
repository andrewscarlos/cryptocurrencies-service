package cache

import (
	"context"
	"cryptocurrencies-service/config"
	"cryptocurrencies-service/entity"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type RedisCache struct {
	db  *redis.Client
	ctx context.Context
}

func NewCacheRedis(ctx context.Context) CacheInterface {
	return &RedisCache{
		db:  config.NewConnectRedis(ctx),
		ctx: ctx,
	}
}

func (c *RedisCache) Get(key string) (entity.Asset, error) {
	val, err := c.db.Get(key).Bytes()
	if err != nil {
		return entity.Asset{}, err
	}
	var asset entity.Asset
	err = json.Unmarshal(val, &asset)
	if err != nil {
		return entity.Asset{}, err
	}
	fmt.Println("Key", key, "Value", &asset)
	return asset, nil
}

func (c *RedisCache) Set(k string, v entity.Asset) error {
	fmt.Println("Key", k, "Value", v)
	val, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.db.Set(k, val, 0).Err()
}

func (c *RedisCache) Delete(key string) error {
	return c.db.Del(key).Err()
}
