package repository

import (
	apperrors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
)

type PaymentRepository interface {
	AddPaymentModeToDB(userPaymentMode *models.UserPaymentMode) *apperrors.AppError
	GetPaymentModeFromDB(userId string) (*models.UserPaymentMode, *apperrors.AppError)
	DBHealthCheck() bool
	UpdatePaymentModeToDB(userPaymentMode *models.UserPaymentMode) *apperrors.AppError
}
