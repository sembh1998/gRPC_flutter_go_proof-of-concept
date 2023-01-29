package model

type MemberMongoModel struct {
	ID                string `bson:"id"`
	UserName          string `bson:"username"`
	FirstName         string `bson:"firstname"`
	LastName          string `bson:"lastname"`
	ProfilePictureURL string `bson:"profilepictureurl"`
}
