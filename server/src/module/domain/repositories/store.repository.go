package repositories

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type StoreRepository interface {
	FindAll() ([]*entities.Store, error)
	FindById(id string) (*entities.Store, error)
	FindByName(name string) ([]*entities.Store, error)
	FindByEstablishmentId(establishmentId string) ([]*entities.Store, error)
}
