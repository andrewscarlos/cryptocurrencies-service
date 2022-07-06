package main

import (
	"cryptocurrencies-service/db"
	"cryptocurrencies-service/pb"
	"cryptocurrencies-service/repository"
	"cryptocurrencies-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	//tronUrl := "https://api.shasta.trongrid.io/wallet/listwitnesses"
	//req, _ := http.NewRequest("GET", tronUrl, nil)
	//res, _ := http.DefaultClient.Do(req)
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

	lis, err := net.Listen("tcp", "localhost:5051")
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
