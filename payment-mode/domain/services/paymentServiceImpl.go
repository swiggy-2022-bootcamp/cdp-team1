package services

import (
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/domain/repository"
	"qwik.in/payment-mode/log"
)

type PaymentServiceImpl struct {
	paymentRepository repository.PaymentRepository
}

func (p PaymentServiceImpl) AddPaymentMode(userPaymentMode *models.UserPaymentMode) error {
	err := p.paymentRepository.AddPaymentModeToDB(userPaymentMode)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p PaymentServiceImpl) GetPaymentMode(userId string) (*models.UserPaymentMode, error) {

	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return userPaymentModes, nil
}

func NewPaymentServiceImpl(paymentRepository repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepository: paymentRepository,
	}
}
