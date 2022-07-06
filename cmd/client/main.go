package main

import (
	"cryptocurrencies-service/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}
	defer connection.Close()
	pb.NewAssetServiceClient(connection)

}