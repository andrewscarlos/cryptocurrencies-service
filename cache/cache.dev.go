package cache

type CacheDev struct {
}

var CacheList = make(map[string]interface{})

func NewCacheDev() CacheInterface {
	return &CacheDev{}
}

func (c *CacheDev) Get(key string) (interface{}, error) {
	return CacheList[key], nil
}
func (c *CacheDev) Set(k string, v interface{}) error {
	CacheList[k] = v
	return nil
}

func (c *CacheDev) Delete(key string) error {
	delete(CacheList, key)
	return nil
}
