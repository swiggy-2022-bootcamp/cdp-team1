package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"qwik.in/payment-mode/mocks"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(paymentRepository *mocks.MockPaymentRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				// Call to MOCK DBHealthCheck() returning true.
				paymentRepository.EXPECT().
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
			name: "InternalServerError",
			buildStubs: func(paymentRepository *mocks.MockPaymentRepository) {
				// Call to MOCK DBHealthCheck() returning false.
				paymentRepository.EXPECT().
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

			//Creating an object og mock repository.
			paymentRepository := mocks.NewMockPaymentRepository(ctrl)
			tc.buildStubs(paymentRepository)

			//Creating handler and setting up router
			healthCheckHandler := NewHealthCheckHandler(paymentRepository)
			server := gin.Default()
			router := server.Group("payment-mode/api")
			router.GET("/", healthCheckHandler.HealthCheck)

			// Making an HTTP call and recording the response
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/payment-mode/api/")
			request := httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestNewHealthCheckHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	paymentRepository := mocks.NewMockPaymentRepository(ctrl)

	healthCheckHandler := NewHealthCheckHandler(paymentRepository)
	assert.Equal(t, healthCheckHandler.paymentRepository, paymentRepository)
}

func requireBodyMatchResposne(t *testing.T, actualResponse *bytes.Buffer, expectedResponse HealthCheckResponse) {
	data, err := ioutil.ReadAll(actualResponse)
	require.NoError(t, err)

	var responseReceived HealthCheckResponse
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, expectedResponse)
}
