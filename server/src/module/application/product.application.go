package application

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/repositories"

type ProductApplication struct {
	ProductRepository repositories.ProductRepository
}

func NewProductApplication(productRepository repositories.ProductRepository) *ProductApplication {
	return &ProductApplication{ProductRepository: productRepository}
}
