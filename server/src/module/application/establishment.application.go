package application

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
)

type EstablishmentApplication struct {
	EstablishmentRepository repositories.EstablishmentRepository
}

func NewEstablishmentApplication(establishmentRepository repositories.EstablishmentRepository) *EstablishmentApplication {
	return &EstablishmentApplication{EstablishmentRepository: establishmentRepository}
}

func (e *EstablishmentApplication) FindAllByStoreId(storeId string) ([]*entities.Establishment, error) {
	return e.EstablishmentRepository.FindAllByStoreId(storeId)
}

func (e *EstablishmentApplication) FindById(id string) (*entities.Establishment, error) {
	return e.EstablishmentRepository.FindById(id)
}
