package repository_test

import (
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/repository"
	mock_repository "cryptocurrencies-service/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestAssetAdapterRepository_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_repository.NewMockAssetRepositoryInterface(ctrl)
	persistence.EXPECT().Insert(gomock.Any()).Return(nil)

	adapter := repository.AssetRepositoryAdapter{
		Persister: persistence,
	}
	err := adapter.Insert(&entity.Asset{
		Id:         bson.NewObjectId(),
		Name:       "John Doe Coin",
		Address:    "foo",
		Blockchain: "bar",
	})
	require.Nil(t, err)
}

func TestAssetRepositoryAdapter_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mock_repository.NewMockAssetRepositoryInterface(ctrl)
	persistence.EXPECT().Delete(gomock.Any()).Return(nil)

	adapter := repository.AssetRepositoryAdapter{
		Persister: persistence,
	}
	err := adapter.Delete("foo")
	require.Nil(t, err)

}

func TestAssetRepositoryAdapter_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var asset entity.Asset
	asset.Id = bson.NewObjectId()
	asset.Value = 1.0
	asset.Name = "foo"
	asset.Address = "bar"
	asset.Blockchain = "baz"
	
	persistence := mock_repository.NewMockAssetRepositoryInterface(ctrl)
	persistence.EXPECT().Read(gomock.Any()).Return(&asset, nil)

	adapter := repository.AssetRepositoryAdapter{
		Persister: persistence,
	}
	result, err := adapter.Read("foo")
	require.Nil(t, err)
	require.Equal(t, &asset, result)
}

func TestAssetRepositoryAdapter_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var asset entity.Asset
	persistence := mock_repository.NewMockAssetRepositoryInterface(ctrl)
	persistence.EXPECT().Update(gomock.Any()).Return(nil)

	adapter := repository.AssetRepositoryAdapter{
		Persister: persistence,
	}
	err := adapter.Update(&asset)
	require.Nil(t, err)
}
