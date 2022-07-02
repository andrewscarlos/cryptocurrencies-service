package service

import (
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
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

func (s *AssetService) Insert(req *pb.Asset) (*pb.Asset, error) {
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
		Id:         req.GetId(),
		Address:    req.GetAddress(),
		Value:      req.GetValue(),
		Name:       req.GetName(),
		Blockchain: req.GetBlockchain(),
	}, nil
}

func (s *AssetService) Read(req *pb.ID) (*pb.Asset, error) {
	result, err := s.AssetRepository.Read(req.GetId())
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

func (s *AssetService) Delete(req *pb.ID) (*pb.ID, error) {
	err := s.AssetRepository.Delete(req.Id)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.ID{
		Id: req.GetId(),
	}, nil
}

func (s *AssetService) Update(req *pb.Asset) (*pb.Asset, error) {
	var assetModel entity.Asset
	asset, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		log.Fatal(err)
	}
	assetModel.Id = asset.Id
	assetModel.Address = asset.Address
	assetModel.Value = asset.Value
	assetModel.Name = asset.Name
	assetModel.Blockchain = asset.Blockchain
	err = s.AssetRepository.Update(&assetModel)
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
