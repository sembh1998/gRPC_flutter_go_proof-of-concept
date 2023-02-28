package controllers

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/application"
	pb "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
)

type Server struct {
	config                   util.Config
	storeApplication         *application.StoreApplication
	memberApplication        *application.MemberApplication
	establishmentApplication *application.EstablishmentApplication
	productApplication       *application.ProductApplication
	pb.UnimplementedModuleServiceServer
}

type ControllerApplicationsEntry struct {
	StoreApplication         *application.StoreApplication
	MemberApplication        *application.MemberApplication
	EstablishmentApplication *application.EstablishmentApplication
	ProductApplication       *application.ProductApplication
}

func NewServer(config util.Config, entry ControllerApplicationsEntry) (*Server, error) {
	return &Server{
		config:                   config,
		storeApplication:         entry.StoreApplication,
		memberApplication:        entry.MemberApplication,
		establishmentApplication: entry.EstablishmentApplication,
		productApplication:       entry.ProductApplication,
	}, nil
}
