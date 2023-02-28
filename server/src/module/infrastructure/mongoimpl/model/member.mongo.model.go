package model

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type MemberMongoModel struct {
	ID                string `bson:"id"`
	UserName          string `bson:"username"`
	FirstName         string `bson:"firstname"`
	LastName          string `bson:"lastname"`
	ProfilePictureURL string `bson:"profilepictureurl"`
}

func (m MemberMongoModel) ToEntity() *entities.Member {
	return &entities.Member{
		ID:                m.ID,
		UserName:          m.UserName,
		FirstName:         m.FirstName,
		LastName:          m.LastName,
		ProfilePictureURL: m.ProfilePictureURL,
	}
}
