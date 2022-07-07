package service

import (
	"context"
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"cryptocurrencies-service/util"
	"gopkg.in/mgo.v2/bson"
	"io"
)

type AssetServiceInterface interface {
	Insert(req *pb.CreateAsset) (*pb.Asset, error)
	Read(req *pb.ID) (*pb.Asset, error)
	Delete(req *pb.ID) (*pb.ID, error)
	Update(req *pb.Asset) (*pb.Asset, error)
	StreamList(stream pb.AssetService_StreamListServer) error
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
	assetModel.Amount = float32(req.GetAmount())

	errInputValidate := validateInputCreate(req)
	if errInputValidate != nil {
		return nil, errInputValidate
	}

	err := s.AssetRepository.Insert(&assetModel)
	if err != nil {
		return nil, util.ErrCreateFailed
	}
	return &pb.Asset{
		Id:         assetModel.Id.Hex(),
		Address:    assetModel.Address,
		Amount:     assetModel.Amount,
		Name:       assetModel.Name,
		Blockchain: assetModel.Blockchain,
	}, nil
}

func (s *AssetService) Read(ctx context.Context, req *pb.ID) (*pb.Asset, error) {
	IsObjectIdHex := bson.IsObjectIdHex(req.GetId())
	if IsObjectIdHex == false {
		return nil, util.ErrInvalidObjectId
	}
	result, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		return nil, util.ErrNotFound
	}
	return &pb.Asset{
		Id:         result.Id.Hex(),
		Address:    result.Address,
		Amount:     float32(result.Amount),
		Name:       result.Name,
		Blockchain: result.Blockchain,
	}, nil

}

func (s *AssetService) Delete(ctx context.Context, req *pb.ID) (*pb.ID, error) {
	IsObjectIdHex := bson.IsObjectIdHex(req.GetId())
	if IsObjectIdHex == false {
		return nil, util.ErrInvalidObjectId
	}
	err := s.AssetRepository.Delete(req.Id)
	if err != nil {
		return nil, util.ErrDeleteFailed
	}
	return &pb.ID{
		Id: req.GetId(),
	}, nil
}

func (s *AssetService) Update(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {
	IsObjectIdHex := bson.IsObjectIdHex(req.GetId())
	if IsObjectIdHex == false {
		return nil, util.ErrInvalidObjectId
	}
	errInputValidate := validateInputUpdate(req)
	if errInputValidate != nil {
		return nil, errInputValidate
	}

	asset, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		return nil, util.ErrNotFound
	}

	asset.Id = bson.ObjectIdHex(req.GetId())
	asset.Address = req.GetAddress()
	asset.Amount = req.GetAmount()
	asset.Name = req.GetName()
	asset.Blockchain = req.GetBlockchain()
	err = s.AssetRepository.Update(asset)
	if err != nil {
		return nil, util.ErrUpdateFailed
	}
	return &pb.Asset{
		Id:         asset.Id.Hex(),
		Address:    asset.Address,
		Amount:     float32(asset.Amount),
		Name:       asset.Name,
		Blockchain: asset.Blockchain,
	}, nil
}

func (s *AssetService) StreamList(stream pb.AssetService_StreamListServer) error {
	assets := []*pb.Asset{}
	for {
		assetRecived, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Assets{
				Assets: assets,
			})
		}
		errInputValidate := validateInputCreate(assetRecived)
		if errInputValidate != nil {
			return util.ErrEmptyInput
		}

		var assetModel entity.Asset
		assetModel.Id = bson.NewObjectId()
		assetModel.Name = assetRecived.GetName()
		assetModel.Address = assetRecived.GetAddress()
		assetModel.Blockchain = assetRecived.GetBlockchain()
		assetModel.Amount = float32(assetRecived.GetAmount())

		err = s.AssetRepository.Insert(&assetModel)

		if err != nil {
			return util.ErrCreateFailed
		}

		assets = append(assets, &pb.Asset{
			Id:         assetModel.Id.Hex(),
			Address:    assetModel.Address,
			Amount:     float32(assetModel.Amount),
			Name:       assetModel.Name,
			Blockchain: assetModel.Blockchain,
		})
		return nil

	}
}

func validateInputUpdate(req *pb.Asset) error {
	if req.GetAddress() == "" || req.GetName() == "" || req.GetBlockchain() == "" || req.GetAmount() == 0 {
		return util.ErrEmptyInput
	}
	return nil
}

func validateInputCreate(req *pb.CreateAsset) error {
	if req.GetAddress() == "" || req.GetName() == "" || req.GetBlockchain() == "" || req.GetAmount() == 0 {
		return util.ErrEmptyInput
	}
	return nil
}
