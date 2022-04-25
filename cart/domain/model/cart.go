package model

type Product struct {
	ProductId string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity"`
}

type Cart struct {
	CustomerId string    `json:"customer_id,omitempty"`
	Products   []Product `json:"products"`
}
