package models

type TransactionDetails struct {
	UserId  string `json:"user_id"`
	OrderId string `json:"order_id" validate:"required"`
	Amount  int    `json:"amount" validate:"required"`
}
