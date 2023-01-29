package mongoimpl

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
)

type EstablishmentMongoInfrastructure struct {
	EstablishmentMongoRepository repositories.EstablishmentRepository
}

func NewEstablishmentMongoInfrastructure(establishmentMongoRepository repositories.EstablishmentRepository) repositories.EstablishmentRepository {
	return &EstablishmentMongoInfrastructure{EstablishmentMongoRepository: establishmentMongoRepository}
}

func (e EstablishmentMongoInfrastructure) FindById(id string) (*entities.Establishment, error) {
	panic("implement me")
}

func (e EstablishmentMongoInfrastructure) FindAllByStoreId(storeId string) ([]*entities.Establishment, error) {
	panic("implement me")
}
