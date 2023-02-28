package controllers

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
	"golang.org/x/net/context"
)

func (s *Server) GetEstablishmentsPerStore(ctx context.Context, req *pb.EstablishmentsByStoreIdRequest) (*pb.EstablishmentsResponse, error) {
	establishments, err := s.establishmentApplication.FindAllByStoreId(req.IdStore)
	if err != nil {
		return &pb.EstablishmentsResponse{
			Error: &pb.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			Success: false,
		}, nil
	}

	var establishmentsResponse []*pb.Establishment
	for _, establishment := range establishments {
		establishmentsResponse = append(establishmentsResponse, &pb.Establishment{
			Id:   establishment.ID,
			Name: establishment.Name,
		})
	}
	return &pb.EstablishmentsResponse{
		Establishments: establishmentsResponse,
		Success:        true,
	}, nil
}

func (s *Server) GetEstablishmentById(ctx context.Context, req *pb.EstablishmentByIdRequest) (*pb.EstablishmentResponse, error) {
	establishment, err := s.establishmentApplication.FindById(req.IdEstablishment)
	if err != nil {
		return &pb.EstablishmentResponse{
			Error: &pb.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			Success: false,
		}, nil
	}

	return &pb.EstablishmentResponse{
		Establishment: &pb.Establishment{
			Id:   establishment.ID,
			Name: establishment.Name,
		},
		Success: true,
	}, nil
}
