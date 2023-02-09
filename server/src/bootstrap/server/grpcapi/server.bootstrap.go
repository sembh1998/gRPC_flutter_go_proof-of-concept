package grpcapi

import (
	"log"
	"net"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/bootstrap/database/connection/mongo"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/application"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/mongoimpl"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/interfaces/grpc/controllers"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func NewServer(config util.Config) (bootstrap.Bootstrap, error) {
	return &Server{}, nil
}

func (s *Server) Initialice(config util.Config) error {

	db, _ := mongo.NewMongoConnection()
	db.Initialice(config)

	conn := mongo.GetMongoConnection().Conn
	storeInfrastructure := mongoimpl.NewStoreMongoInfrastructure(conn)
	if storeInfrastructure == nil {
		log.Fatal("cannot create store infrastructure")
	}
	storeApplication := application.NewStoreApplication(storeInfrastructure)

	controllerEntry := controllers.ControllerApplicationsEntry{
		StoreApplication: storeApplication,
	}

	server, err := controllers.NewServer(config, controllerEntry)
	if err != nil {
		log.Fatal("cannot creating grpc server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterModuleServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Println("grpc server running on: ", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("cannot serve grpc server: ", err)
	}
	return nil
}
