package main

import (
	"log"
	"net"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap/server/grpcapi"
	pb "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

}

func runGrpcServer(config util.Config) {
	server, err := grpcapi.NewServer(config)
	if err != nil {
		log.Fatal("cannot creating grpc server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStoreServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Println("grpc server running on: ", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("cannot serve grpc server: ", err)
	}
}
