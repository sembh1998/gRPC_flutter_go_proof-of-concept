package mongoimpl

import (
	"context"
	"log"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/mongoimpl/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EstablishmentMongoInfrastructure struct {
	DB *mongo.Database
}

func NewEstablishmentMongoInfrastructure(db *mongo.Database) repositories.EstablishmentRepository {
	return &EstablishmentMongoInfrastructure{DB: db}
}

func (e EstablishmentMongoInfrastructure) FindById(id string) (*entities.Establishment, error) {
	var store *model.StoreMongoModel
	coll := e.DB.Collection("stores")

	filter := bson.M{"establishments": bson.M{"$elemMatch": bson.M{"id": id}}}

	// create a projection that excludes the products field
	projection := bson.M{"establishments": bson.M{
		"id":   1,
		"name": 1,
	}}

	err := coll.FindOne(context.Background(), filter, &options.FindOneOptions{
		Projection: projection,
	}).Decode(&store)
	if err != nil {
		return nil, err
	}
	log.Printf("store: %v", store)
	log.Printf("number of establishments: %v", len(store.Establishments))
	for _, establishment := range store.Establishments {
		log.Printf("establishment: %v", establishment)
		if establishment.ID == id {
			log.Printf("found")
			return establishment.ToEntity(), nil
		}
	}
	return nil, nil
}

func (e EstablishmentMongoInfrastructure) FindAllByStoreId(storeId string) ([]*entities.Establishment, error) {
	var store *model.StoreMongoModel
	coll := e.DB.Collection("stores")
	// create a projection that excludes the products field
	projection := bson.M{"establishments": bson.M{
		"id":   1,
		"name": 1,
	}}
	err := coll.FindOne(context.Background(), bson.M{"id": storeId}, &options.FindOneOptions{
		Projection: projection,
	}).Decode(&store)
	if err != nil {
		return nil, err
	}
	establishments := make([]*entities.Establishment, 0)
	for _, establishment := range store.Establishments {
		log.Printf("establishment: %v", establishment)
		establishments = append(establishments, establishment.ToEntity())
	}
	return establishments, nil
}
