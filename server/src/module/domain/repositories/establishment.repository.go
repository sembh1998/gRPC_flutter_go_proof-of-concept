package repositories

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type EstablishmentRepository interface {
	FindById(id string) (*entities.Establishment, error)
	FindAllByStoreId(storeId string) ([]*entities.Establishment, error)
}
