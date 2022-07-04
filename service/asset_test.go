package service_test

import (
	"cryptocurrencies-service/pb"
	mock_service "cryptocurrencies-service/service/mocks"
	"cryptocurrencies-service/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssetService_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var req pb.CreateAsset
	var res pb.Asset

	req.Value = 1.0
	req.Name = "foo"
	req.Address = "bar"
	req.Blockchain = "baz"

	res.Id = "62bf4284956789b5c6ea0edb"
	res.Value = 1.0
	res.Name = "foo"
	res.Address = "bar"
	res.Blockchain = "baz"

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&res, nil)

	result, err := services.Insert(&req)
	require.Nil(t, err)
	require.Equal(t, &res, result)
}

func TestAssetService_Insert_WhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var res pb.Asset
	var req pb.CreateAsset
	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&res, util.ErrCreateFailed)
	_, err := services.Insert(&req)
	require.Equal(t, "asset created failed", err.Error())
}

func TestAssetService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var id pb.ID

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Delete(gomock.Any()).Return(&id, nil)

	result, err := services.Delete(&id)
	require.Nil(t, err)
	require.Equal(t, &id, result)
}

func TestAssetService_Read(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var id pb.ID
	var asset pb.Asset

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Read(gomock.Any()).Return(&asset, nil)

	result, err := services.Read(&id)
	require.Nil(t, err)
	require.Equal(t, &asset, result)
}

func TestAssetService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var req pb.Asset

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Update(gomock.Any()).Return(&req, nil)

	result, err := services.Update(&req)
	require.Nil(t, err)
	require.Equal(t, &req, result)
}
