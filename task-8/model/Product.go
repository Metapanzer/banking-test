package model

type Product struct {
	Sku             string  `json:"sku"`
	ProductName     string  `json:"productName"`
	QuantityInStock int     `json:"quantityInStock"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
}
