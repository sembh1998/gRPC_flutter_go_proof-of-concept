package model

import "github.com/sembh1998/gRPC_flutter_go_proof-of-concept/server/src/module/domain/entities"

type ProductMongoModel struct {
	ID              string  `bson:"id"`
	Name            string  `bson:"name"`
	SKU             string  `bson:"sku"`
	Barcode         string  `bson:"barcode"`
	AbbreviatedName string  `bson:"abbreviated_name"`
	ProductImageURL string  `bson:"product_image_url"`
	Price           float64 `bson:"price"`
	Stock           int64   `bson:"stock"`
}

func (p ProductMongoModel) ToEntity() *entities.Product {
	return &entities.Product{
		ID:              p.ID,
		Name:            p.Name,
		SKU:             p.SKU,
		Barcode:         p.Barcode,
		AbbreviatedName: p.AbbreviatedName,
		ProductImageURL: p.ProductImageURL,
		Price:           p.Price,
		Stock:           p.Stock,
	}
}
