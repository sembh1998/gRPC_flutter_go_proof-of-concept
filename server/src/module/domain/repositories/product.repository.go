package repositories

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type ProductRepository interface {
	FindAllByEstablishmentId(establishmentId string) ([]*entities.Product, error)
	FindById(id string) (*entities.Product, error)
}
