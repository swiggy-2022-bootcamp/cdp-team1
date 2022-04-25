package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	app_erros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/mocks"
	"testing"
)

func TestTransactionHandler_AddTransactionPoints(t *testing.T) {
	transactionAmount := &models.TransactionDetails{
		Amount:  20000,
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459tuf",
		OrderId: "OA-123",
	}
	transactionAmountValidationError := &models.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459tuf",
		OrderId: "OA-123",
	}
	testCases := []struct {
		name          string
		buildStubs    func(transactionService *mocks.MockTransactionService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "FailureWithMalformedBody",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					AddTransactionPoints(transactionAmount).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "FailureWithValidationError",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					AddTransactionPoints(transactionAmount).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "FailureWithStatusNotFound",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					AddTransactionPoints(transactionAmount).
					Times(1).
					Return(app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureWithStatusInternalServerError",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					AddTransactionPoints(transactionAmount).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Success",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					AddTransactionPoints(transactionAmount).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(transactionAmount)
			require.NoError(t, err)

			transactionService := mocks.NewMockTransactionService(ctrl)
			tc.buildStubs(transactionService)

			server := NewServer(transactionService)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/transaction/api/transaction/%s", transactionAmount.UserId)
			var request *http.Request
			if tc.name == "FailureWithMalformedBody" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "FailureWithValidationError" {
				data, err := json.Marshal(transactionAmountValidationError)
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

func TestTransactionHandler_GetTransactionPointsByUserID(t *testing.T) {
	userId := "bb912edc-50d9-42d7-b7a1-9ce66d459tuf"
	transactionPoints := 100

	testCases := []struct {
		name          string
		buildStubs    func(transactionService *mocks.MockTransactionService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					GetTransactionPointsByUserId(userId).
					Times(1).
					Return(transactionPoints, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "FailureWithUserNotFound",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					GetTransactionPointsByUserId(userId).
					Times(1).
					Return(-1, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureWithInternalServerError",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					GetTransactionPointsByUserId(userId).
					Times(1).
					Return(-1, app_erros.NewUnexpectedError(""))
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

			transactionService := mocks.NewMockTransactionService(ctrl)
			tc.buildStubs(transactionService)

			server := NewServer(transactionService)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/transaction/api/transaction/%s", userId)
			var request *http.Request
			request = httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestTransactionHandler_UseTransactionPoints(t *testing.T) {
	transactionAmount := &models.TransactionDetails{
		Amount:  20000,
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459tuf",
		OrderId: "OA-123",
	}
	transactionAmountValidationError := &models.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459tuf",
		OrderId: "OA-123",
	}
	testCases := []struct {
		name          string
		buildStubs    func(transactionService *mocks.MockTransactionService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "FailureWithMalformedBody",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(0).
					Return(false, nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "FailureWithValidationError",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(0).
					Return(false, nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "FailureWithUserNotFound",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(1).
					Return(false, transactionAmount, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureWithStatusInternalServerError",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(1).
					Return(false, transactionAmount, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "FailureWithExpectationFailed",
			buildStubs: func(transactionService *mocks.MockTransactionService) {
				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(1).
					Return(false, transactionAmount, app_erros.NewExpectationFailed(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusExpectationFailed, recorder.Code)
			},
		},
		{
			name: "Success",
			buildStubs: func(transactionService *mocks.MockTransactionService) {

				transactionService.EXPECT().
					UseTransactionPoints(transactionAmount).
					Times(1).
					Return(true, transactionAmount, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(transactionAmount)
			require.NoError(t, err)

			transactionService := mocks.NewMockTransactionService(ctrl)
			tc.buildStubs(transactionService)

			server := NewServer(transactionService)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/transaction/api/transaction/use-transaction-points/%s", transactionAmount.UserId)
			var request *http.Request
			if tc.name == "FailureWithMalformedBody" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "FailureWithValidationError" {
				data, err := json.Marshal(transactionAmountValidationError)
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
