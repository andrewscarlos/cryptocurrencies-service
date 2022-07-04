package service

import (
	"context"
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AssetServiceInterface interface {
	Insert(req *pb.Asset) (*pb.Asset, error)
	Read(req *pb.ID) (*pb.Asset, error)
	Delete(req *pb.ID) (*pb.ID, error)
	Update(req *pb.Asset) (*pb.Asset, error)
}

type AssetService struct {
	pb.UnimplementedAssetServiceServer
	AssetRepository repository.AssetRepositoryInterface
}

func NewAssetService(assetRepository repository.AssetRepositoryInterface) *AssetService {
	return &AssetService{

		AssetRepository: assetRepository,
	}
}

func (s *AssetService) Insert(ctx context.Context, req *pb.CreateAsset) (*pb.Asset, error) {
	var assetModel entity.Asset
	assetModel.Id = bson.NewObjectId()
	assetModel.Name = req.GetName()
	assetModel.Address = req.GetAddress()
	assetModel.Blockchain = req.GetBlockchain()
	assetModel.Value = float32(req.GetValue())

	err := s.AssetRepository.Insert(&assetModel)
	if err != nil {
		return nil, err
		//log.Fatalf("Could not Insert request: %v", err)
	}
	//TODO validate request body fields
	return &pb.Asset{
		Id:         assetModel.Id.Hex(),
		Address:    assetModel.Address,
		Value:      assetModel.Value,
		Name:       assetModel.Name,
		Blockchain: assetModel.Blockchain,
	}, nil
}

func (s *AssetService) Read(ctx context.Context, req *pb.ID) (*pb.Asset, error) {
	result, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		return nil, err
	}
	fmt.Println("result", result)
	return &pb.Asset{
		Id:         result.Id.Hex(),
		Address:    result.Address,
		Value:      float32(result.Value),
		Name:       result.Name,
		Blockchain: result.Blockchain,
	}, nil

}

func (s *AssetService) Delete(ctx context.Context, req *pb.ID) (*pb.ID, error) {
	err := s.AssetRepository.Delete(req.Id)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.ID{
		Id: req.GetId(),
	}, nil
}

func (s *AssetService) Update(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {

	asset, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		log.Fatal(err)
	}

	asset.Id = bson.ObjectIdHex(req.GetId())
	asset.Address = req.GetAddress()
	asset.Value = req.GetValue()
	asset.Name = req.GetName()
	asset.Blockchain = req.GetBlockchain()
	err = s.AssetRepository.Update(asset)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Asset{
		Id:         asset.Id.Hex(),
		Address:    asset.Address,
		Value:      float32(asset.Value),
		Name:       asset.Name,
		Blockchain: asset.Blockchain,
	}, nil
}
