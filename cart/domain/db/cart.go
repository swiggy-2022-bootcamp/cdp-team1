package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductId primitive.ObjectID `json:"id,omitempty"`
	Quantity  int                `json:"quantity"`
}

type Cart struct {
	Id         string             `json:"id,omitempty"`
	CustomerId primitive.ObjectID `json:"customer_id,omitempty"`
	Products   []Product          `json:"products"`
}
