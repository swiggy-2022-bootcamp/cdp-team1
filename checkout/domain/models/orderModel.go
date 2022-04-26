package models

//OrderProductModel - Order Service.
type OrderProductModel struct {
	ProductID string `json:"product_id,omitempty" dynamodbav:"product_id"`
	Quantity  int    `json:"quantity" dynamodb:"quantity"`
}

//OrderModel - Order Service.
type OrderModel struct {
	OrderID    string              `json:"id,omitempty" dynamodbav:"id"`
	CustomerID string              `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	Status     string              `json:"status" dynamodbav:"status"`
	Datetime   string              `json:"datetime" dynamodbav:"datetime"`
	Products   []OrderProductModel `json:"orders" dynamodbav:"orders"`
	Invoice    string              `json:"invoice" dynamodbav:"invoice"`
}
