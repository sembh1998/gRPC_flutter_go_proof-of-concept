package model

type EstablishmentMongoModel struct {
	ID       string              `bson:"id"`
	Name     string              `bson:"name"`
	Products []ProductMongoModel `bson:"products"`
}
