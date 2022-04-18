package services

import (
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
)

type TransactionService interface {
	AddTransactionPoints(transactionAmount *models.TransactionDetails) *apperros.AppError
	GetTransactionPointsByUserId(userId string) (int, *apperros.AppError)
	CalculateTransactionPoints(transactionAmount *models.TransactionDetails) int
	UseTransactionPoints(transactionAmount *models.TransactionDetails) (bool, *models.TransactionDetails, *apperros.AppError)
	UpdateTransactionPoints(transactionPoint int, userId string) *apperros.AppError
}
