package main

import (
	"context"
	"cryptocurrencies-service/cache"
	"cryptocurrencies-service/config"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"cryptocurrencies-service/service"
	"cryptocurrencies-service/util"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", os.Getenv("GRPC_HOST"))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	grpcServer := grpc.NewServer()

	dbConn, err := config.NewConnection()
	if err != nil {
		log.Fatalf(util.ErrNotConnecInDatabase.Error())
	}
	defer dbConn.Close()

	assetMongoRepository := repository.NewAssetRepository(dbConn)
	assetRepository := repository.NewAssetRepositoryAdapter(assetMongoRepository)
	cacheRedis := cache.NewCacheRedis(context.Background())
	cache := cache.NewCache(cacheRedis)
	pb.RegisterAssetServiceServer(grpcServer, service.NewAssetService(assetRepository, cache))
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
