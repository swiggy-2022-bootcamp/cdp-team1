package repository

import (
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
)

type TransactionRepository interface {
	AddTransactionPointsFromDB(transaction *models.Transaction) *apperros.AppError
	GetTransactionPointsByUserIdFromDB(userId string) (int, *apperros.AppError)
	DBHealthCheck() bool
	UpdateTransactionPointsToDB(transaction *models.Transaction) *apperros.AppError
}
