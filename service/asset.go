package service

import (
	"context"
	"cryptocurrencies-service/cache"
	"cryptocurrencies-service/entity"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"cryptocurrencies-service/util"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
)

type AssetServiceInterface interface {
	Insert(req *pb.CreateAsset) (*pb.Asset, error)
	Read(req *pb.ID) (*pb.Asset, error)
	Delete(req *pb.ID) (*pb.ID, error)
	Update(req *pb.Asset) (*pb.Asset, error)
	StreamList(stream pb.AssetService_StreamListServer) error
	GetAll(void *pb.Void) (*pb.Assets, error)
}

type AssetService struct {
	pb.UnimplementedAssetServiceServer
	AssetRepository repository.AssetRepositoryInterface
	Cache           cache.CacheInterface
}

func NewAssetService(assetRepository repository.AssetRepositoryInterface, cache cache.CacheInterface) *AssetService {
	return &AssetService{
		AssetRepository: assetRepository,
		Cache:           cache,
	}
}

func (s *AssetService) Insert(ctx context.Context, req *pb.CreateAsset) (*pb.Asset, error) {
	errInputValidate := validateInputCreate(req)
	if errInputValidate != nil {
		return nil, errInputValidate
	}
	assetModel := entity.Asset{
		Id:         bson.NewObjectId(),
		Name:       req.GetName(),
		Address:    req.GetAddress(),
		Blockchain: req.GetBlockchain(),
		Amount:     float32(req.GetAmount()),
	}

	err := s.AssetRepository.Insert(&assetModel)
	if err != nil {
		return nil, util.ErrCreateFailed
	}
	s.Cache.Set(assetModel.Id.Hex(), assetModel)
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
	assetCache, err := s.Cache.Get(req.GetId())
	if err == nil {
		assetReturn := &pb.Asset{
			Id:         assetCache.Id.Hex(),
			Address:    assetCache.Address,
			Amount:     float32(assetCache.Amount),
			Name:       assetCache.Name,
			Blockchain: assetCache.Blockchain,
		}
		return assetReturn, nil
	}
	result, err := s.AssetRepository.Read(req.GetId())
	if err != nil {
		return nil, util.ErrNotFound
	}
	err = s.addAssetToCacheIfNotExists(result)
	if err != nil {
		log.Println(err)
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
	s.Cache.Delete(req.Id)
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
	s.Cache.Set(asset.Id.Hex(), entity.Asset{
		Id:         asset.Id,
		Address:    asset.Address,
		Amount:     asset.Amount,
		Name:       asset.Name,
		Blockchain: asset.Blockchain,
	})

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

		errInputValidate := validateInputCreate(assetRecived)
		if errInputValidate != nil {
			return util.ErrEmptyInput
		}

		if err == io.EOF {
			return stream.SendAndClose(&pb.Assets{
				Assets: assets,
			})
		}

		assetModel := entity.Asset{
			Id:         bson.NewObjectId(),
			Name:       assetRecived.GetName(),
			Address:    assetRecived.GetAddress(),
			Blockchain: assetRecived.GetBlockchain(),
			Amount:     float32(assetRecived.GetAmount()),
		}

		err = s.AssetRepository.Insert(&assetModel)

		if err != nil {
			return util.ErrCreateFailed
		}
		s.Cache.Set(assetModel.Id.Hex(), assetModel)
		assets = append(assets, &pb.Asset{
			Id:         assetModel.Id.Hex(),
			Address:    assetModel.Address,
			Amount:     float32(assetModel.Amount),
			Name:       assetModel.Name,
			Blockchain: assetModel.Blockchain,
		})
	}
}

func (s *AssetService) GetAll(ctx context.Context, void *pb.Void) (*pb.Assets, error) {
	getAllassets, err := s.AssetRepository.GetAll()
	if err != nil {
		return nil, util.ErrEmptyAssetList
	}
	var assetList []*pb.Asset

	for _, asset := range getAllassets {
		assetList = append(assetList, &pb.Asset{
			Id:         asset.Id.Hex(),
			Address:    asset.Address,
			Amount:     float32(asset.Amount),
			Name:       asset.Name,
			Blockchain: asset.Blockchain,
		})
	}
	return &pb.Assets{
		Assets: assetList,
	}, nil
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

func (s *AssetService) addAssetToCacheIfNotExists(asset *entity.Asset) error {
	return s.Cache.Set(asset.Id.Hex(), entity.Asset{
		Id:         asset.Id,
		Address:    asset.Address,
		Amount:     float32(asset.Amount),
		Name:       asset.Name,
		Blockchain: asset.Blockchain,
	})
}
