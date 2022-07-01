package service

import (
	"context"
	"cryptocurrencies-service/model"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/reposiroty"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AssetService struct {
	pb.UnimplementedAssetServiceServer
	assetRepository reposiroty.AssetRepositoryInterface
}

func NewAssetService(assetRepository reposiroty.AssetRepositoryInterface) *AssetService {
	return &AssetService{
		assetRepository: assetRepository,
	}
}

func (s *AssetService) Insert(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {
	var assetModel model.Asset
	//if req.Id != "" {
	//	req.Id = bson.NewObjectId().String()
	//}
	assetModel.Id = bson.NewObjectId()
	assetModel.Name = req.GetName()
	assetModel.Address = req.GetAddress()
	assetModel.Blockchain = req.GetBlockchain()
	assetModel.Value = float32(req.GetValue())

	err := s.assetRepository.Insert(&assetModel)
	if err != nil {
		log.Fatalf("Could not Insert request: %v", err)
	}
	return &pb.Asset{
		Id:         req.GetId(),
		Address:    req.GetAddress(),
		Value:      req.GetValue(),
		Name:       req.GetName(),
		Blockchain: req.GetBlockchain(),
	}, nil
}

func (s *AssetService) Read(ctx context.Context, req *pb.ID) (*pb.Asset, error) {
	result, err := s.assetRepository.Read(req.GetId())
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Asset{
		Id:         string(result.Id),
		Address:    result.Address,
		Value:      float32(result.Value),
		Name:       result.Name,
		Blockchain: result.Blockchain,
	}, nil

}

func (s *AssetService) Delete(ctx context.Context, req *pb.ID) (*pb.ID, error) {
	err := s.assetRepository.Delete(req.Id)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.ID{
		Id: req.GetId(),
	}, nil
}

func (s *AssetService) Update(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {
	var assetModel model.Asset
	asset, err := s.assetRepository.Read(req.GetId())
	if err != nil {
		log.Fatal(err)
	}
	assetModel.Id = asset.Id
	assetModel.Address = asset.Address
	assetModel.Value = asset.Value
	assetModel.Name = asset.Name
	assetModel.Blockchain = asset.Blockchain
	err = s.assetRepository.Update(&assetModel)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Asset{
		Id:         string(assetModel.Id),
		Address:    assetModel.Address,
		Value:      float32(assetModel.Value),
		Name:       assetModel.Name,
		Blockchain: assetModel.Blockchain,
	}, nil
}
