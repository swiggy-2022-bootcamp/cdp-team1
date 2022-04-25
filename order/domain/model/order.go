package model

type Product struct {
	ProductID string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity"`
}

type Order struct {
	OrderId    string    `json:"order_id,omitempty"`
	CustomerId string    `json:"customer_id,omitempty"`
	Status     string    `json:"status"`
	Datetime   string    `json:"datetime"`
	Products   []Product `json:"orders"`
	Invoice    string    `json:"invoice"`
	Amount     int       `json:"amount"`
}
