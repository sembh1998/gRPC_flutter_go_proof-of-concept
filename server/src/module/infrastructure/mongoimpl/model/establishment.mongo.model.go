package model

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type EstablishmentMongoModel struct {
	ID       string              `bson:"id"`
	Name     string              `bson:"name"`
	Products []ProductMongoModel `bson:"products"`
}

func (e EstablishmentMongoModel) ToEntity() *entities.Establishment {
	return &entities.Establishment{
		ID:   e.ID,
		Name: e.Name,
	}
}
