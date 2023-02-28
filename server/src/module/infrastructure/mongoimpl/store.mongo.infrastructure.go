package mongoimpl

import (
	"context"

	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/infrastructure/mongoimpl/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoreMongoInfrastructure struct {
	DB *mongo.Database
}

func NewStoreMongoInfrastructure(db *mongo.Database) repositories.StoreRepository {
	return &StoreMongoInfrastructure{DB: db}
}

func (s *StoreMongoInfrastructure) FindAll() ([]*entities.Store, error) {
	var stores []*model.StoreMongoModel
	coll := s.DB.Collection("stores")
	cursor, err := coll.Find(context.Background(), bson.D{}, &options.FindOptions{
		Sort: bson.M{
			"name": 1,
		},
		Projection: bson.M{
			"establishments": 0,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &stores); err != nil {
		return nil, err
	}
	var storesEntity []*entities.Store
	for _, store := range stores {
		storesEntity = append(storesEntity, store.ToEntity())
	}
	return storesEntity, nil
}

func (s *StoreMongoInfrastructure) FindById(id string) (*entities.Store, error) {
	var store *model.StoreMongoModel
	coll := s.DB.Collection("stores")
	err := coll.FindOne(context.Background(), bson.M{"id": id}, &options.FindOneOptions{
		Projection: bson.M{
			"establishments": 0,
		},
	}).Decode(&store)
	if err != nil {
		return nil, err
	}
	return store.ToEntity(), nil
}

func (s *StoreMongoInfrastructure) FindByName(name string) ([]*entities.Store, error) {
	var stores []*model.StoreMongoModel
	coll := s.DB.Collection("stores")
	cursor, err := coll.Find(context.Background(), bson.M{"name": name}, &options.FindOptions{
		Projection: bson.M{
			"establishments": 0,
		},
	})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &stores); err != nil {
		return nil, err
	}
	var storesEntity []*entities.Store
	for _, store := range stores {
		storesEntity = append(storesEntity, store.ToEntity())
	}
	return storesEntity, nil
}
