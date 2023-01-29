package application

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"

type EstablishmentApplication struct {
	EstablishmentRepository repositories.EstablishmentRepository
}

func NewEstablishmentApplication(establishmentRepository repositories.EstablishmentRepository) *EstablishmentApplication {
	return &EstablishmentApplication{EstablishmentRepository: establishmentRepository}
}
