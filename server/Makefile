proto: 
	rm -f ./src/module/infrastructure/serializer/grpc/pb/*.pb.go
	protoc -I ../protos \
	--go_out=./src/module/infrastructure/serializer/grpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=./src/module/infrastructure/serializer/grpc/pb --go-grpc_opt=paths=source_relative \
	../protos/*.proto

.PHONY: proto