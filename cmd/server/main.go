package main

import (
	"cryptocurrencies-service/db"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/reposiroty"
	"cryptocurrencies-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	grpcServer := grpc.NewServer()

	dbConn := db.NewConnection()
	defer dbConn.Close()

	assetRepository := reposiroty.NewAssetRepository(dbConn)

	pb.RegisterAssetServiceServer(grpcServer, service.NewAssetService(assetRepository))
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
