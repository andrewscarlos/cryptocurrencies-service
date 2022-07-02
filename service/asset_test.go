package service_test

import (
	"cryptocurrencies-service/pb"
	mock_service "cryptocurrencies-service/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssetService_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var req pb.Asset

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&req, nil)

	result, err := services.Insert(&req)
	require.Nil(t, err)
	require.Equal(t, &req, result)
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
