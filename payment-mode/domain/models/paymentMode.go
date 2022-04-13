package models

type PaymentMode struct {
	Mode       string `json:"mode" dynamodbav:"mode" validate:"required"`
	CardNumber int    `json:"cardNumber" dynamodbav:"card_number"`
	Balance    int    `json:"balance" dynamodbav:"balance"`
}
