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
		buildStubs    func(paymentService *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "BadRequestFailure",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(0).
					Return(nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			paymentRepository := mocks.NewMockPaymentRepository(ctrl)
			tc.buildStubs(paymentRepository)

			server := NewServer(paymentRepository)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/paymentmethods/%s", userId)
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
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
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessUpdatePaymentMode",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				paymentRepository.EXPECT().
					GetPaymentModeFromDB(userId).
					Times(1).
					Return(&userPaymentMode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchPaymentMode(t, recorder.Body, userPaymentMode)
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
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
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

			paymentRepository := mocks.NewMockPaymentRepository(ctrl)
			tc.buildStubs(paymentRepository)

			server := NewServer(paymentRepository)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/paymentmethods/%s", userId)
			request := httptest.NewRequest(http.MethodGet, url, nil)

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
