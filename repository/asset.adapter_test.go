package repository_test

import (
	"cryptocurrencies-service/entity"
	mock_model "cryptocurrencies-service/entity/mocks"
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

	mock_model.NewMockAssetInterface(ctrl)
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
