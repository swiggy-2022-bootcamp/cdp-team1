package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/domain/services"
	"qwik.in/payment-mode/mocks"
	"testing"
)

func TestPaymentHandler_AddPaymentMode(t *testing.T) {

	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	paymentModeValidationError := models.PaymentMode{
		CardNumber: 4242424242424242,
		Balance:    500,
	}
	testCases := []struct {
		name          string
		buildStubs    func(paymentService *mocks.MockPaymentService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			// Missing payment mode body in the request
			name: "BadRequestFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					AddPaymentMode(&paymentMode, userId).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			// Required field not present in request body.
			name: "ValidationFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					AddPaymentMode(&paymentMode, userId).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			// Payment added successfully.
			name: "Success",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					AddPaymentMode(&paymentMode, userId).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			// Payment mode addition failed with Internal server error.
			name: "Failure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					AddPaymentMode(&paymentMode, userId).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(paymentMode)
			require.NoError(t, err)

			// Mock payment service
			paymentService := mocks.NewMockPaymentService(ctrl)
			tc.buildStubs(paymentService)

			server := NewServer(paymentService)

			// Making an HTTP request
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/paymentmethods/%s", userId)
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "ValidationFailure" {
				data, err := json.Marshal(paymentModeValidationError)
				require.NoError(t, err)
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			} else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
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
		buildStubs    func(paymentService *mocks.MockPaymentService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessGetPaymentMode",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					GetPaymentMode(userId).
					Times(1).
					Return(&userPaymentMode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentMode(t, recorder.Body, userPaymentMode)
			},
		},
		{
			name: "FailureUserNotFound",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					GetPaymentMode(userId).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					GetPaymentMode(userId).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			paymentService := mocks.NewMockPaymentService(ctrl)
			tc.buildStubs(paymentService)

			server := NewServer(paymentService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/paymentmethods/%s", userId)
			request := httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestPaymentHandler_SetPaymentMode(t *testing.T) {
	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	paymentModeValidationError := models.PaymentMode{
		CardNumber: 4242424242424242,
		Balance:    500,
	}
	testCases := []struct {
		name          string
		buildStubs    func(paymentService *mocks.MockPaymentService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "BadRequestFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					SetPaymentMode(&paymentMode, userId).
					Times(0).
					Return(false, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ValidationFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					SetPaymentMode(&paymentMode, userId).
					Times(0).
					Return(false, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Success",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					SetPaymentMode(userId, paymentMode).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Failure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					SetPaymentMode(userId, paymentMode).
					Times(1).
					Return(false, app_errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(paymentMode)
			require.NoError(t, err)

			paymentService := mocks.NewMockPaymentService(ctrl)
			tc.buildStubs(paymentService)

			server := NewServer(paymentService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/setpaymentmethods/%s", userId)
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "ValidationFailure" {
				data, err := json.Marshal(paymentModeValidationError)
				require.NoError(t, err)
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			} else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestPaymentHandler_CompletePayment(t *testing.T) {
	userId := "abcd-efgh-1234-4321"
	paymentMode := models.PaymentMode{
		Mode:       "Credit Card",
		CardNumber: 4242424242424242,
	}
	paymentRequest := models.PaymentRequest{
		SelectedPaymentMode: paymentMode,
		UserId:              userId,
		OrderId:             "OA-123",
		OrderAmount:         1000,
	}
	paymentRequestValidationError := models.PaymentRequest{
		SelectedPaymentMode: paymentMode,
		UserId:              userId,
	}
	testCases := []struct {
		name          string
		buildStubs    func(paymentService *mocks.MockPaymentService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "BadRequestFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					CheckBalanceAndCompletePayment(&paymentRequest).
					Times(0).
					Return(false, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ValidationFailure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					CheckBalanceAndCompletePayment(&paymentRequest).
					Times(0).
					Return(false, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Success",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					CheckBalanceAndCompletePayment(&paymentRequest).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
		{
			name: "Failure",
			buildStubs: func(paymentService *mocks.MockPaymentService) {
				paymentService.EXPECT().
					CheckBalanceAndCompletePayment(&paymentRequest).
					Times(1).
					Return(false, app_errors.NewRequestNotAcceptedError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotAcceptable, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(paymentRequest)
			require.NoError(t, err)

			paymentService := mocks.NewMockPaymentService(ctrl)
			tc.buildStubs(paymentService)

			server := NewServer(paymentService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/pay")
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "ValidationFailure" {
				data, err := json.Marshal(paymentRequestValidationError)
				require.NoError(t, err)
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			} else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestNewPaymentHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	paymentRepository := mocks.NewMockPaymentRepository(ctrl)

	paymentService := services.NewPaymentServiceImpl(paymentRepository)
	paymentHandler := NewPaymentHandler(paymentService)

	assert.Equal(t, paymentHandler.paymentService, paymentService)
}

func requireBodyMatchPaymentMode(t *testing.T, body *bytes.Buffer, requiredResponse models.UserPaymentMode) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var responseReceived models.UserPaymentMode
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, requiredResponse)
}
