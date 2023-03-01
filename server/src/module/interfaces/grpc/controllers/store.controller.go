package controllers

import (
	"context"
	"log"

	pb "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/serializer/grpc/pb"
)

func (s *Server) GetStores(ctx context.Context, req *pb.StoreRequest) (*pb.StoresResponse, error) {
	stores, err := s.storeApplication.FindAll()
	if err != nil {
		log.Printf("Error: %v", err)
		return &pb.StoresResponse{
			Error: &pb.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			Success: false,
		}, nil
	}

	var storesResponse []*pb.Store
	for _, store := range stores {
		storesResponse = append(storesResponse, &pb.Store{
			Id:   store.ID,
			Name: store.Name,
		})
	}
	return &pb.StoresResponse{
		Stores:  storesResponse,
		Success: true,
	}, nil

}

func (s *Server) GetStoreById(ctx context.Context, req *pb.StoreByIdRequest) (*pb.StoreResponse, error) {
	store, err := s.storeApplication.FindById(req.Id)
	if err != nil {
		log.Printf("Error: %v", err)
		return &pb.StoreResponse{
			Error: &pb.Error{
				Code:    500,
				Message: "Internal Server Error",
			},
			Success: false,
		}, nil
	}

	return &pb.StoreResponse{
		Store: &pb.Store{
			Id:   store.ID,
			Name: store.Name,
		},
		Success: true,
	}, nil

}

func (s *Server) GetStoreByName(context.Context, *pb.StoreByNameRequest) (*pb.StoreResponse, error) {
	panic("implement me")
}
