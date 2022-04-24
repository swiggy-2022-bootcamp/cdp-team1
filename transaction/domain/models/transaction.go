package models

type Transaction struct {
	UserId            string `json:"user_id" dynamodbav:"user_id" validate:"required"`
	TransactionPoints int    `json:"transaction_points" dynamodbav:"transaction_points"`
}
