package services

import (
	"context"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/repository"
	"qwik.in/payment-mode/log"
	"qwik.in/payment-mode/protos"
)

var paymentRepository repository.PaymentRepository

type PaymentProtoServer struct {
	protos.UnimplementedPaymentServer
}

func NewPaymentProtoService(pr repository.PaymentRepository) PaymentProtoServer {
	paymentRepository = pr
	return PaymentProtoServer{}
}

func (s PaymentProtoServer) CompletePayment(ctx context.Context, paymentRequest *protos.PaymentRequest) (*protos.PaymentResponse, error) {
	paymentResponse := &protos.PaymentResponse{
		IsPaymentSuccessful: false,
	}
	userPaymentModes, err := paymentRepository.GetPaymentModeFromDB(paymentRequest.GetUserId())

	if err != nil {
		return paymentResponse, err.Error()
	}

	for i, availablePaymentMode := range userPaymentModes.PaymentModes {

		if availablePaymentMode.Mode == paymentRequest.GetPaymentMode().Mode && availablePaymentMode.CardNumber == int(paymentRequest.GetPaymentMode().CardNumber) {

			if availablePaymentMode.Balance >= int(paymentRequest.GetAmount()) {

				availablePaymentMode.Balance -= int(paymentRequest.GetAmount())
				userPaymentModes.PaymentModes[i] = availablePaymentMode
				updateErr := paymentRepository.UpdatePaymentModeToDB(userPaymentModes)

				if updateErr != nil {
					return paymentResponse, updateErr.Error()
				} else {
					paymentResponse.IsPaymentSuccessful = true
					return paymentResponse, nil
				}

			} else {
				return paymentResponse, app_errors.NewRequestNotAcceptedError("Insufficient funds, payment failed.").Error()
			}
		}
	}

	return paymentResponse, app_errors.NewNotFoundError("Payment method is not added for the current user.").Error()
}

func (s PaymentProtoServer) GetPaymentModes(ctx context.Context, paymentModeRequest *protos.PaymentModeRequest) (*protos.PaymentModeResponse, error) {

	paymentModes := make([]*protos.PaymentMode, 0, 0)
	response := &protos.PaymentModeResponse{
		UserId:       paymentModeRequest.GetUserId(),
		PaymentModes: paymentModes,
	}

	// Fetch payment modes for the given user
	userPaymentModes, err := paymentRepository.GetPaymentModeFromDB(paymentModeRequest.GetUserId())
	if err != nil {
		log.Error(err)
		return nil, err.Error()
	}

	//Convert dynamoDB object into proto message object.
	var paymentModeProto *protos.PaymentMode
	for _, paymentMode := range userPaymentModes.PaymentModes {
		paymentModeProto = &protos.PaymentMode{
			Mode:       paymentMode.Mode,
			CardNumber: int64(paymentMode.CardNumber),
			Balance:    int32(paymentMode.Balance),
		}
		paymentModes = append(paymentModes, paymentModeProto)
	}
	response.PaymentModes = paymentModes
	return response, nil
}
