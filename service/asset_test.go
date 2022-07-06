package service_test

import (
	"cryptocurrencies-service/pb"
	mock_service "cryptocurrencies-service/service/mocks"
	"cryptocurrencies-service/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssetServiceInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var req pb.CreateAsset
	var res pb.Asset

	req.Amount = 1.0
	req.Name = "foo"
	req.Address = "bar"
	req.Blockchain = "baz"

	res.Id = "62bf4284956789b5c6ea0edb"
	res.Amount = 1.0
	res.Name = "foo"
	res.Address = "bar"
	res.Blockchain = "baz"

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&res, nil)

	result, err := services.Insert(&req)
	require.Nil(t, err)
	require.Equal(t, &res, result)
}

func TestAssetServiceInsertWhenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var res pb.Asset
	var req pb.CreateAsset
	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&res, util.ErrCreateFailed)
	_, err := services.Insert(&req)
	require.Equal(t, "asset created failed", err.Error())

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Insert(gomock.Any()).Return(&res, util.ErrEmptyInput)

	_, err = services.Insert(&req)
	require.Equal(t, "empty input", err.Error())

}

func TestAssetServiceDeleteAllCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var id pb.ID

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Delete(gomock.Any()).Return(&id, nil)
	result, err := services.Delete(&id)
	require.Nil(t, err)
	require.Equal(t, &id, result)

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Delete(gomock.Any()).Return(nil, util.ErrInvalidObjectId)
	_, err = services.Delete(&id)
	require.Equal(t, "invalid objectId", err.Error())

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Delete(gomock.Any()).Return(nil, util.ErrDeleteFailed)
	_, err = services.Delete(&id)
	require.Equal(t, "asset deleted failed", err.Error())

}

func TestAssetServiceRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var id pb.ID
	var asset pb.Asset

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Read(gomock.Any()).Return(&asset, nil)

	result, err := services.Read(&id)
	require.Nil(t, err)
	require.Equal(t, &asset, result)

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Read(gomock.Any()).Return(nil, util.ErrInvalidObjectId)
	_, err = services.Read(&id)
	require.Equal(t, "invalid objectId", err.Error())

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Read(gomock.Any()).Return(nil, util.ErrNotFound)
	_, err = services.Read(&id)
	require.Equal(t, "asset not found", err.Error())

}

func TestAssetServiceUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var req pb.Asset

	req.Id = "62bf4284956789b5c6ea0edb"
	req.Name = "foo_updated"
	req.Address = "bar_updated"
	req.Blockchain = "baz_updated"
	req.Amount = 2.0

	services := mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Update(gomock.Any()).Return(&req, nil)

	result, err := services.Update(&req)
	require.Nil(t, err)
	require.Equal(t, &req, result)

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Update(gomock.Any()).Return(nil, util.ErrInvalidObjectId)
	_, err = services.Update(&req)
	require.Equal(t, "invalid objectId", err.Error())

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Update(gomock.Any()).Return(nil, util.ErrNotFound)
	_, err = services.Update(&req)
	require.Equal(t, "asset not found", err.Error())

	services = mock_service.NewMockAssetServiceInterface(ctrl)
	services.EXPECT().Update(gomock.Any()).Return(nil, util.ErrEmptyInput)
	_, err = services.Update(&req)
	require.Equal(t, "empty input", err.Error())

}
