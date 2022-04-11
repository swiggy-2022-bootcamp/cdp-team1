package models

type PaymentMode struct {
	Mode       string `json:"mode" bson:"mode" validate:"required"`
	CardNumber int    `json:"cardNumber" bson:"card_number"`
	Balance    int    `json:"balance" bson:"balance"`
}
