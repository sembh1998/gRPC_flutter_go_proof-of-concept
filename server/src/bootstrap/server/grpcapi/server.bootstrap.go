package grpcapi

import (
	pb "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/util"
)

type Server struct {
	config util.Config
	pb.UnimplementedStoreServiceServer
}

func NewServer(config util.Config) (*Server, error) {
	return &Server{config: config}, nil
}
