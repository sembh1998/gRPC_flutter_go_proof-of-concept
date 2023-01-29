package model

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
