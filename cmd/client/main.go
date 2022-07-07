package main

import (
	"cryptocurrencies-service/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection, err := grpc.Dial(os.Getenv("GRPC_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}
	defer connection.Close()
	pb.NewAssetServiceClient(connection)

}
