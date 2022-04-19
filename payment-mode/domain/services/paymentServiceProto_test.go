package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/mocks"
	"qwik.in/payment-mode/protos"
	"reflect"
	"testing"
)

func TestPaymentProtoServer_CompletePayment(t *testing.T) {
	userId := "abcd-efgh-1234-4321"
	//To be used as selected Payment mode and will also be returned by the database.
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
		Balance:    10000,
	}

	//To be used as the payment mode that will be updated to the database after successful payment
	updatedPaymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
		Balance:    5000,
	}

	// the payment request made by the user.
	paymentRequest := &protos.PaymentRequest{
		PaymentMode: &protos.PaymentMode{
			Mode:       "Credit Card",
			CardNumber: 4242424242424242,
			Balance:    10000,
		},
		Amount:  5000,
		OrderId: "OA-101",
		UserId:  userId,
	}

	currentPaymentMode := make([]models.PaymentMode, 0, 0)
	userPaymentModeSuccess := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(currentPaymentMode, paymentMode),
	}

	updatedPaymentModeToDB := make([]models.PaymentMode, 0, 0)
	userPaymentModeUpdated := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(updatedPaymentModeToDB, updatedPaymentMode),
	}

	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, isPaymentSuccessful bool, errors error)
	}{
		{
			name: "FailureGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isPaymentModeValid bool, errors error) {
				require.Equal(t, false, isPaymentModeValid)
				require.Equal(t, app_errors.NewUnexpectedError("").Error(), errors)
			},
		},
		{
			name: "PaymentSuccessful",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeSuccess, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeUpdated).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors error) {
				require.Equal(t, true, isPaymentSuccessful)
				require.NoError(t, errors)
			},
		},
		{
			name: "PaymentUpdationFailure",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentMode.Balance = 10000
				currentPaymentMode := make([]models.PaymentMode, 0, 0)
				userPaymentModeSuccess := models.UserPaymentMode{
					UserId:       userId,
					PaymentModes: append(currentPaymentMode, paymentMode),
				}
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeSuccess, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeUpdated).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors error) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewUnexpectedError("").Error(), errors)
			},
		},
		{
			name: "PaymentFailed",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				//Requested amount is greater than balance amount
				paymentMode.Balance = 2000
				currentPaymentMode := make([]models.PaymentMode, 0, 0)
				userPaymentModeSuccess := models.UserPaymentMode{
					UserId:       userId,
					PaymentModes: append(currentPaymentMode, paymentMode),
				}

				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeSuccess, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeUpdated).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors error) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewRequestNotAcceptedError("Insufficient funds, payment failed.").Error(), errors)
			},
		},
		{
			name: "PaymentModeInvalid",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				//Requested payment mode is not available for the user.

				currentPaymentMode := make([]models.PaymentMode, 0, 0)
				userPaymentModeSuccess := models.UserPaymentMode{
					UserId:       userId,
					PaymentModes: currentPaymentMode,
				}

				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeSuccess, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeUpdated).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors error) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewNotFoundError("Payment method is not added for the current user.").Error(), errors)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			paymentRepository := mocks.NewMockPaymentRepository(ctrl)
			tc.buildStubs(paymentRepository)
			paymentProtoService := NewPaymentProtoService(paymentRepository)
			isPaymentSuccessful, err := paymentProtoService.CompletePayment(context.Background(), paymentRequest)
			tc.checkResponse(t, isPaymentSuccessful.IsPaymentSuccessful, err)
		})
	}
}

func TestPaymentProtoServer_GetPaymentModes(t *testing.T) {
	userId := "abcd-efgh-1234-4321"

	//Object that will be returned from the database.
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	currentPaymentMode := make([]models.PaymentMode, 0, 0)
	userPaymentMode := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(currentPaymentMode, paymentMode),
	}

	//Proto Object that will be returned to the client
	paymentRequest := &protos.PaymentModeRequest{
		UserId: userId,
	}
	paymentModeProto := &protos.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	currentPaymentModeProto := make([]*protos.PaymentMode, 0, 0)
	paymentModeResponse := &protos.PaymentModeResponse{
		UserId:       userId,
		PaymentModes: append(currentPaymentModeProto, paymentModeProto),
	}

	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, response *protos.PaymentModeResponse, err error)
	}{
		{
			name: "SuccessGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentMode, nil)
			},
			checkResponse: func(t *testing.T, response *protos.PaymentModeResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, response, paymentModeResponse)

			},
		},
		{
			name: "FailureUserPaymentModeNotFound",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, response *protos.PaymentModeResponse, err error) {
				require.Equal(t, true, reflect.ValueOf(response).IsNil())
				require.Equal(t, app_errors.NewNotFoundError("").Error(), err)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T, response *protos.PaymentModeResponse, err error) {
				require.Equal(t, true, reflect.ValueOf(response).IsNil())
				require.Equal(t, app_errors.NewUnexpectedError("").Error(), err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			paymentRepository := mocks.NewMockPaymentRepository(ctrl)
			tc.buildStubs(paymentRepository)
			paymentProtoServer := NewPaymentProtoService(paymentRepository)
			result, err := paymentProtoServer.GetPaymentModes(context.Background(), paymentRequest)
			tc.checkResponse(t, result, err)
		})
	}
}
