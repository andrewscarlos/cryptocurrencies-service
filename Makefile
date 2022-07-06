

grpcui:
	grpcui -plaintext localhost:5051

server:
	go run cmd/server/main.go

proto:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

mock:
	 mockgen -destination=service/mocks/asset.go -source=service/asset.go service