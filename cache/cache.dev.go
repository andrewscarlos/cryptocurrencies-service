package cache

import "cryptocurrencies-service/entity"

type CacheDev struct {
}

var CacheList = make(map[string]entity.Asset)

func NewCacheDev() CacheInterface {
	return &CacheDev{}
}

func (c *CacheDev) Get(key string) (entity.Asset, error) {
	return CacheList[key], nil
}
func (c *CacheDev) Set(k string, v entity.Asset) error {
	CacheList[k] = v
	return nil
}

func (c *CacheDev) Delete(key string) error {
	delete(CacheList, key)
	return nil
}
