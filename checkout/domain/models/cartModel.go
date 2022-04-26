package models

//CartModel - Cart Service
type CartModel struct {
	ID               string `json:"id,omitempty" dynamodbav:"id"`
	CustomerID       string `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	CartProductModel []struct {
		ProductID string `json:"product_id,omitempty" dynamodbav:"product_id"`
		Quantity  int    `json:"quantity" dynamodbav:"quantity"`
	} `json:"products" dynamodbav:"products"`
}
