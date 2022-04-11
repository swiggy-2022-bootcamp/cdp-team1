package services

import "qwik.in/payment-mode/domain/models"

type PaymentService interface {
	AddPaymentMode(userPaymentMode *models.UserPaymentMode) error
	GetPaymentMode(userId string) (*models.UserPaymentMode, error)
}
