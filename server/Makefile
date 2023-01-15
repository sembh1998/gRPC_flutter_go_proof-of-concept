GOOS = $(shell uname -s)
GOARCH = $(shell uname -m)

build:
	go build -o bin/server_$(GOOS)-$(GOARCH) ./main.go

proto: 
	rm -f ./src/module/infrastructure/serializer/grpc/pb/*.pb.go
	protoc -I ../protos \
	--go_out=./src/module/infrastructure/serializer/grpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=./src/module/infrastructure/serializer/grpc/pb --go-grpc_opt=paths=source_relative \
	../protos/*.proto

.PHONY: proto build