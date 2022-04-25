package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"qwik.in/transaction/mocks"
	"testing"
)

func TestHealthCheckHandler_HealthCheck(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(transactionRepository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			//Database connection is up.
			name: "Success",
			buildStubs: func(transactionRepository *mocks.MockTransactionRepository) {
				transactionRepository.EXPECT().
					DBHealthCheck().
					Times(1).
					Return(true)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchResposne(t, recorder.Body, HealthCheckResponse{Server: "Server is up",
					Database: "Database is up"})
			},
		},
		{
			// Database connection is down
			name: "InternalServerError",
			buildStubs: func(transactionRepository *mocks.MockTransactionRepository) {
				transactionRepository.EXPECT().
					DBHealthCheck().
					Times(1).
					Return(false)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				requireBodyMatchResposne(t, recorder.Body, HealthCheckResponse{Server: "Server is up",
					Database: "Database is down"})
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// TransactionRepository mock object
			transactionRepository := mocks.NewMockTransactionRepository(ctrl)
			tc.buildStubs(transactionRepository)
			healthCheckHandler := NewHealthCheckHandler(transactionRepository)
			server := gin.Default()
			router := server.Group("transaction/api")

			router.GET("/", healthCheckHandler.HealthCheck)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/transaction/api/")
			request := httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchResposne(t *testing.T, body *bytes.Buffer, requiredResponse HealthCheckResponse) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var responseReceived HealthCheckResponse
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, requiredResponse)
}
