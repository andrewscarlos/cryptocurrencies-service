package service

import (
	"context"
	"cryptocurrencies-service/pb"
)

type AssetService struct {
	pb.UnimplementedAssetServiceServer
}

//rpc Insert(Asset) returns (Asset){};
//rpc Read(ID) returns (Asset){}
//rpc Delete(ID) returns (ID){}
//rpc Update(Asset) returns (Asset){}

//string Id = 1;
//string Address = 2;
//float  Value = 3;
//string Name = 4;
//string Blockchain = 5;

func NewAssetService() *AssetService {
	return &AssetService{}
}

func (*AssetService) Insert(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {
	return &pb.Asset{
		Id:         "123",
		Address:    req.GetAddress(),
		Value:      req.GetValue(),
		Name:       req.GetName(),
		Blockchain: req.GetBlockchain(),
	}, nil
}

func (*AssetService) Read(ctx context.Context, req *pb.ID) (*pb.Asset, error) {
	return nil, nil
}

func (*AssetService) Delete(ctx context.Context, req *pb.ID) (*pb.ID, error) {
	return nil, nil
}

func (*AssetService) Update(ctx context.Context, req *pb.Asset) (*pb.Asset, error) {
	return nil, nil
}
