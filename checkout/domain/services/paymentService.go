package services

import (
	"qwik.in/checkout/domain/repository"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
)

//PaymentService ..
type PaymentService interface {
	GetDefaultPaymentMode() (*repository.PaymentUserPaymentMode, *errs.AppError)
}

//PaymentServiceImpl ..
type PaymentServiceImpl struct {
	paymentRepo repository.PaymentRepo
}

//PaymentServiceFunc ..
func PaymentServiceFunc(paymentRepository repository.PaymentRepo) PaymentService {
	return &PaymentServiceImpl{
		paymentRepo: paymentRepository,
	}
}

//GetDefaultPaymentMode ..
func (p PaymentServiceImpl) GetDefaultPaymentMode() (*repository.PaymentUserPaymentMode, *errs.AppError) {
	data, err := p.paymentRepo.GetPaymentDetailsImpl()
	if err != nil {
		logger.Error("Error on paymentService.go -> GetDefaultPaymentMode() Function !", err)
		return nil, err
	}
	return data, nil
}
