package mongoimpl

import (
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"
	"github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberMongoInfrastructure struct {
	DB *mongo.Database
}

func NewMemberMongoInfrastructure(db *mongo.Database) repositories.MemberRepository {
	return &MemberMongoInfrastructure{DB: db}
}

func (m *MemberMongoInfrastructure) FindAll() ([]*entities.Member, error) {
	panic("implement me")
}

func (m *MemberMongoInfrastructure) FindById(id string) (*entities.Member, error) {
	panic("implement me")
}

func (m *MemberMongoInfrastructure) Create(member *entities.Member) (*entities.Member, error) {
	panic("implement me")
}
