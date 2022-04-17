package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/mocks"
	"reflect"
	"testing"
)

func TestPaymentHandler_AddPaymentMode(t *testing.T) {

	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	userPaymentModeUpdate := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: make([]models.PaymentMode, 0, 0),
	}
	currentPaymentMode := make([]models.PaymentMode, 0, 0)
	userPaymentModeConflict := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(currentPaymentMode, paymentMode),
	}

	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, expected interface{}, actual interface{})
	}{
		{
			name: "ConflictPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeConflict, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeConflict).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, app_errors.NewConflictRequestError("This payment mode is already present in your account."), actual)
			},
		},
		{
			name: "SuccessUpdatePaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeUpdate, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeConflict).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
			},
		},
		{
			name: "FailureUpdatePaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeUpdate, nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeConflict).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, app_errors.NewUnexpectedError(""), expected)
			},
		},
		{
			name: "SuccessAddPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

				paymentRepository.EXPECT().
					AddPaymentModeToDB(&userPaymentModeConflict).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, true, reflect.ValueOf(expected).IsNil())
			},
		},
		{
			name: "FailureAddPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

				paymentRepository.EXPECT().
					AddPaymentModeToDB(&userPaymentModeConflict).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, app_errors.NewUnexpectedError(""), expected)
			},
		},
		{
			name: "FailureGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))

				paymentRepository.EXPECT().
					AddPaymentModeToDB(&userPaymentModeConflict).
					Times(0).
					Return(nil)

				paymentRepository.EXPECT().
					UpdatePaymentModeToDB(&userPaymentModeConflict).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, app_errors.NewUnexpectedError(""), expected)
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

			paymentServiceImpl := NewPaymentServiceImpl(paymentRepository)

			err := paymentServiceImpl.AddPaymentMode(&paymentMode, userId)
			tc.checkResponse(t, nil, err)
		})
	}
}

func TestPaymentHandler_GetPaymentMode(t *testing.T) {
	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	currentPaymentMode := make([]models.PaymentMode, 0, 0)
	userPaymentMode := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(currentPaymentMode, paymentMode),
	}

	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, paymentModeFromDB interface{}, err interface{})
	}{
		{
			name: "SuccessGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentMode, nil)
			},
			checkResponse: func(t *testing.T, paymentModeFromDB interface{}, err interface{}) {
				require.Equal(t, true, reflect.ValueOf(err).IsNil())
				require.Equal(t, paymentModeFromDB, &userPaymentMode)

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
			checkResponse: func(t *testing.T, paymentModeFromDB interface{}, err interface{}) {
				require.Equal(t, true, reflect.ValueOf(paymentModeFromDB).IsNil())
				require.Equal(t, app_errors.NewNotFoundError(""), err)
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
			checkResponse: func(t *testing.T, paymentModeFromDB interface{}, err interface{}) {
				require.Equal(t, true, reflect.ValueOf(paymentModeFromDB).IsNil())
				require.Equal(t, app_errors.NewUnexpectedError(""), err)
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
			paymentServiceImpl := NewPaymentServiceImpl(paymentRepository)
			result, err := paymentServiceImpl.GetPaymentMode(userId)
			tc.checkResponse(t, result, err)
		})
	}
}

func TestPaymentServiceImpl_SetPaymentMode(t *testing.T) {
	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	currentPaymentMode := make([]models.PaymentMode, 0, 0)
	userPaymentModeSuccess := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: append(currentPaymentMode, paymentMode),
	}
	userPaymentModeFailure := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: currentPaymentMode,
	}
	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, isPaymentModeValid bool, errors interface{})
	}{
		{
			name: "FailureGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isPaymentModeValid bool, errors interface{}) {
				require.Equal(t, false, isPaymentModeValid)
				require.Equal(t, app_errors.NewUnexpectedError(""), errors)
			},
		},
		{
			name: "SuccessSetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeSuccess, nil)
			},
			checkResponse: func(t *testing.T, isPaymentModeValid bool, errors interface{}) {
				require.Equal(t, true, isPaymentModeValid)
				require.Equal(t, true, reflect.ValueOf(errors).IsNil())
			},
		},
		{
			name: "FailureSetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentModeFailure, nil)
			},
			checkResponse: func(t *testing.T, isPaymentModeValid bool, errors interface{}) {
				require.Equal(t, false, isPaymentModeValid)
				require.Equal(t, app_errors.NewNotFoundError("The given payment mode doesn't exist for the current user."), errors)
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
			paymentServiceImpl := NewPaymentServiceImpl(paymentRepository)
			result, err := paymentServiceImpl.SetPaymentMode(userId, paymentMode)
			tc.checkResponse(t, result, err)
		})
	}
}

func TestPaymentServiceImpl_CheckBalanceAndCompletePayment(t *testing.T) {
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
	paymentRequest := models.PaymentRequest{
		SelectedPaymentMode: paymentMode,
		OrderAmount:         5000,
		OrderId:             "OA-101",
		UserId:              userId,
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
		checkResponse func(t *testing.T, isPaymentSuccessful bool, errors interface{})
	}{
		{
			name: "FailureGetPaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isPaymentModeValid bool, errors interface{}) {
				require.Equal(t, false, isPaymentModeValid)
				require.Equal(t, app_errors.NewUnexpectedError(""), errors)
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
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors interface{}) {
				require.Equal(t, true, isPaymentSuccessful)
				require.Equal(t, true, reflect.ValueOf(errors).IsNil())
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
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors interface{}) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewUnexpectedError(""), errors)
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
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors interface{}) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewRequestNotAcceptedError("Insufficient funds, payment failed."), errors)
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
			checkResponse: func(t *testing.T, isPaymentSuccessful bool, errors interface{}) {
				require.Equal(t, false, isPaymentSuccessful)
				require.Equal(t, app_errors.NewNotFoundError("Payment method is not added for the current user."), errors)
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
			paymentServiceImpl := NewPaymentServiceImpl(paymentRepository)
			isPaymentSuccessful, err := paymentServiceImpl.CheckBalanceAndCompletePayment(&paymentRequest)
			tc.checkResponse(t, isPaymentSuccessful, err)
		})
	}
}
