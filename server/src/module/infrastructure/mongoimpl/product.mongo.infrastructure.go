package mongoimpl

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductMongoInfrastructure struct {
	DB *mongo.Database
}

func NewProductMongoInfrastructure(db *mongo.Database) repositories.ProductRepository {
	return &ProductMongoInfrastructure{DB: db}
}

func (p *ProductMongoInfrastructure) FindAllByEstablishmentId(establishmentId string) ([]*entities.Product, error) {
	panic("implement me")
}

func (p *ProductMongoInfrastructure) FindById(id string) (*entities.Product, error) {
	panic("implement me")
}
