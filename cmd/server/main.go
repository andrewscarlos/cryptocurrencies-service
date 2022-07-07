package main

import (
	"cryptocurrencies-service/db"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"cryptocurrencies-service/service"
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

	dbConn := db.NewConnection()
	defer dbConn.Close()

	assetMongoRepository := repository.NewAssetRepository(dbConn)
	assetRepository := repository.NewAssetRepositoryAdapter(assetMongoRepository)

	pb.RegisterAssetServiceServer(grpcServer, service.NewAssetService(assetRepository))
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}

//func getUrl() string {
//	return fmt.Sprintf(os.Getenv("GRPC_HOST"),os.Getenv("GRPC_PORT"))
//}
