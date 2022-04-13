package repository

import "qwik.in/payment-mode/domain/models"

type PaymentRepository interface {
	AddPaymentModeToDB(userPaymentMode *models.UserPaymentMode) error
	GetPaymentModeFromDB(userId string) (*models.UserPaymentMode, error)
	DBHealthCheck() bool
}
