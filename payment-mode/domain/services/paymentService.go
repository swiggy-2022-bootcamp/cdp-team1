package services

import (
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
)

type PaymentService interface {
	AddPaymentMode(paymentMode *models.PaymentMode, userId string) *app_errors.AppError
	GetPaymentMode(userId string) (*models.UserPaymentMode, *app_errors.AppError)
	GetUserId(token string) string
	SetPaymentMode(userId string, paymentMode models.PaymentMode) (bool, *app_errors.AppError)
	CheckBalanceAndCompletePayment(paymentRequest *models.PaymentRequest) (bool, *app_errors.AppError)
}
