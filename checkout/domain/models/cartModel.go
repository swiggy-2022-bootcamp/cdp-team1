package models

//CartModel - Cart Service
type CartModel struct {
	Id               string `json:"id,omitempty" dynamodbav:"id"`
	CustomerId       string `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	CartProductModel []struct {
		ProductId string `json:"product_id,omitempty" dynamodbav:"product_id"`
		Quantity  int    `json:"quantity" dynamodbav:"quantity"`
	} `json:"products" dynamodbav:"products"`
}
