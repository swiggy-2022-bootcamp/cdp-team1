package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductID primitive.ObjectID `json:"product_id,omitempty"`
	Quantity  int                `json:"quantity"`
}

type Order struct {
	OrderId    string             `json:"id,omitempty"`
	CustomerId primitive.ObjectID `json:"customer_id,omitempty"`
	Status     string             `json:"status"`
	Datetime   string             `json:"datetime"`
	Products   []Product          `json:"orders"`
}
