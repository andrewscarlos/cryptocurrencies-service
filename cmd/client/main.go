package main

import (
	"context"
	"cryptocurrencies-service/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server %v", err)
	}
	defer connection.Close()
	client := pb.NewAssetServiceClient(connection)
	Insert(client)
}

func Insert(client pb.AssetServiceClient) {
	req := &pb.Asset{
		Id:         "0123456789ab0123456789ab", //bson.NewObjectId().String()
		Address:    "HNLQ4sDKQFcDJegPwnwd3N39TBdRQW9Sfd",
		Value:      50000,
		Name:       "Jhon doe Coin",
		Blockchain: "foo",
	}
	res, err := client.Insert(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}
