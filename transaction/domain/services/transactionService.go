package services

import (
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
)

type TransactionService interface {
	AddTransactionPoints(transactionAmount *models.TransactionAmount) *apperros.AppError
	GetTransactionPointsByUserId(userId string) (int, *apperros.AppError)
	CalculateTransactionPoints(transactionAmount *models.TransactionAmount) int
	UseTransactionPoints(transactionAmount models.TransactionAmount) (bool, *models.TransactionAmount)
	UpdateTransactionPoints(transactionPoint int, userId string) *apperros.AppError
}
