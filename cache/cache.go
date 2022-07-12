package cache

import "cryptocurrencies-service/entity"

type Key string
type Value interface{}

type CacheInterface interface {
	Get(key string) (entity.Asset, error)
	Set(key string, value entity.Asset) error
	Delete(key string) error
}

func NewCache(cp CacheInterface) CacheInterface {
	return &Cache{
		CatcherPersistence: cp,
	}
}

type Cache struct {
	CatcherPersistence CacheInterface
}

func (c *Cache) Get(key string) (entity.Asset, error) {
	return c.CatcherPersistence.Get(key)
}

func (c *Cache) Set(key string, value entity.Asset) error {
	return c.CatcherPersistence.Set(key, value)
}

func (c *Cache) Delete(key string) error {
	return c.CatcherPersistence.Delete(key)
}
