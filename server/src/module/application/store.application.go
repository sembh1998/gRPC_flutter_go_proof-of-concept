package application

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
)

type StoreApplication struct {
	StoreRepository repositories.StoreRepository
}

func NewStoreApplication(storeRepository repositories.StoreRepository) *StoreApplication {
	return &StoreApplication{StoreRepository: storeRepository}
}

func (s *StoreApplication) FindAll() ([]*entities.Store, error) {
	return s.StoreRepository.FindAll()
}
