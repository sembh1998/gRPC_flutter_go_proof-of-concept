package application

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"

type MemberApplication struct {
	MemberRepository repositories.MemberRepository
}

func NewMemberApplication(memberRepository repositories.MemberRepository) *MemberApplication {
	return &MemberApplication{MemberRepository: memberRepository}
}
