package cache

import (
	"cryptocurrencies-service/entity"
	"errors"
)

type CacheDev struct {
}

var CacheList = make(map[string]entity.Asset)

func NewCacheDev() CacheInterface {
	return &CacheDev{}
}

func (c *CacheDev) Get(key string) (entity.Asset, error) {
	item, ok := CacheList[key]
	if !ok {
		return entity.Asset{}, errors.New("item not found")
	}
	return item, nil
}
func (c *CacheDev) Set(k string, v entity.Asset) error {
	CacheList[k] = v
	return nil
}

func (c *CacheDev) Delete(key string) error {
	delete(CacheList, key)
	return nil
}
