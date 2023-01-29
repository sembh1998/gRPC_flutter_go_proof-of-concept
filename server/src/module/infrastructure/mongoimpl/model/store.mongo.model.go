package model

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type StoreMongoModel struct {
	ID             string                    `bson:"id"`
	Name           string                    `bson:"name"`
	Establishments []EstablishmentMongoModel `bson:"establishments"`
}

func (m *StoreMongoModel) ToEntity() *entities.Store {
	return &entities.Store{
		ID:   m.ID,
		Name: m.Name,
	}
}
