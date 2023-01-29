package repositories

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type MemberRepository interface {
	FindAll() ([]*entities.Member, error)
	FindById(id string) (*entities.Member, error)
	Create(member *entities.Member) (*entities.Member, error)
}
