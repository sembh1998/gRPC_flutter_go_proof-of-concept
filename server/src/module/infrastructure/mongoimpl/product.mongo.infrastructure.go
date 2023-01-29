package mongoimpl

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
)

type ProductMongoInfrastructure struct {
	ProductRepository repositories.ProductRepository
}

func NewProductMongoInfrastructure(productRepository repositories.ProductRepository) repositories.ProductRepository {
	return &ProductMongoInfrastructure{ProductRepository: productRepository}
}

func (p *ProductMongoInfrastructure) FindAllByEstablishmentId(establishmentId string) ([]*entities.Product, error) {
	panic("implement me")
}

func (p *ProductMongoInfrastructure) FindById(id string) (*entities.Product, error) {
	panic("implement me")
}
