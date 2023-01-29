package entities

type Product struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	SKU             string  `json:"sku"`
	Barcode         string  `json:"barcode"`
	AbbreviatedName string  `json:"abbreviated_name"`
	ProductImageURL string  `json:"product_image_url"`
	Price           float64 `json:"price"`
	Stock           int64   `json:"stock"`
}
