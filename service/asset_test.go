package service_test

import (
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/repository"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var conn repository.AssetRepositoryInterface

//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}

func getAssetModel() *entity.Asset {
	var asset entity.Asset
	asset.Id = bson.NewObjectId()
	asset.Name = "John Doe Coin"
	asset.Address = "foo"
	asset.Blockchain = "bar"
	asset.Value = 1.0
	return &asset
}
func TestAssetService_Insert(t *testing.T) {
	err := conn.Insert(getAssetModel())
	require.Equal(t, "", err.Error())
}
