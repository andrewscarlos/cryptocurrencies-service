package cache_test

import (
	"cryptocurrencies-service/cache"
	"cryptocurrencies-service/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	cachedev := cache.NewCacheDev()
	cache := cache.NewCache(cachedev)
	assetModel := entity.Asset{
		Id:         bson.NewObjectId(),
		Name:       "John Doe Coin",
		Address:    "foo",
		Blockchain: "bar",
		Amount:     1.0,
	}
	err := cache.Set("foo", assetModel)
	require.Nil(t, err)

	asset, err := cache.Get("foo")
	require.Nil(t, err)
	assert.Equal(t, assetModel, asset)
}

func TestCacheDelete(t *testing.T) {
	cachedev := cache.NewCacheDev()
	cache := cache.NewCache(cachedev)
	assetModel := entity.Asset{
		Id:         bson.NewObjectId(),
		Name:       "John Doe Coin",
		Address:    "foo",
		Blockchain: "bar",
		Amount:     1.0,
	}
	cache.Set("foo", assetModel)
	err := cache.Delete("foo")
	require.Nil(t, err)
}
