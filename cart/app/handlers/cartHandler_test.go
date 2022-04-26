package handlers

// import (
// 	"bytes"
// 	"cartService/domain/model"
// 	"cartService/internal/error"
// 	mocks "cartService/mock"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDUiLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE2NTE3Njk3ODl9.49_x8PNd8j2TH_TsLZLyFIiw9DUFU3-LplAYx-uU3UM"

// func TestCarthandler_CreateCart(t *testing.T) {

// 	customer_id := "1234"
// 	products := model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	}

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	invalid_cart := model.Cart{
// 		Products: []model.Product{products},
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartService *mocks.MockCartService)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "SuccessCreateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken(token).
// 					Times(1).
// 					Return("", nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusAccepted, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FailureCreateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken("").
// 					Times(1).
// 					Return("", &error.AppError{})
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FailureCreateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken(token).
// 					Times(1).
// 					Return("", nil)

// 				cartService.EXPECT().
// 					AddToCart(&invalid_cart, customer_id).
// 					Times(1).
// 					Return(&error.AppError{})
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			data, err := json.Marshal(old_cart)
// 			require.NoError(t, err)

// 			// Mock payment service
// 			cartService := mocks.NewMockCartService(ctrl)
// 			tc.buildStubs(cartService)

// 			server := NewServer(cartService)

// 			// Making an HTTP request
// 			recorder := httptest.NewRecorder()
// 			url := fmt.Sprintf("/api/paymentmethods")
// 			var request *http.Request
// 			if tc.name == "BadRequestFailure" {
// 				request = httptest.NewRequest(http.MethodPost, url, nil)
// 			} else if tc.name == "ValidationFailure" {
// 				data, err := json.Marshal(paymentModeValidationError)
// 				require.NoError(t, err)
// 				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			} else {
// 				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			}
// 			request.Header.Set("Authorization", token)
// 			server.ServeHTTP(recorder, request)

// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func TestCarthandler_UpdateCart(t *testing.T) {

// 	customer_id := "1234"
// 	product_id := "1"
// 	quantity := 1

// 	products := model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	}

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	// invalid_cart := model.Cart{
// 	// 	Products: []model.Product{products},
// 	// }

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartService *mocks.MockCartService)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "SuccessUpdateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken(token).
// 					Times(1).
// 					Return("", nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusAccepted, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FailureUpdateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken("").
// 					Times(1).
// 					Return("", &error.AppError{})
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FailureUpdateCart",
// 			buildStubs: func(cartService *mocks.MockCartService) {
// 				cartService.EXPECT().
// 					UserIDFromAuthToken(token).
// 					Times(1).
// 					Return("", nil)

// 				cartService.EXPECT().
// 					UpdateCart(customer_id, product_id, quantity).
// 					Times(1).
// 					Return(&error.AppError{})
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			data, err := json.Marshal(old_cart)
// 			require.NoError(t, err)

// 			// Mock payment service
// 			cartService := mocks.NewMockCartService(ctrl)
// 			tc.buildStubs(cartService)

// 			server := NewServer(cartService)

// 			// Making an HTTP request
// 			recorder := httptest.NewRecorder()
// 			url := fmt.Sprintf("/api/paymentmethods")
// 			var request *http.Request
// 			if tc.name == "BadRequestFailure" {
// 				request = httptest.NewRequest(http.MethodPost, url, nil)
// 			} else if tc.name == "ValidationFailure" {
// 				data, err := json.Marshal(paymentModeValidationError)
// 				require.NoError(t, err)
// 				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			} else {
// 				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			}
// 			request.Header.Set("Authorization", token)
// 			server.ServeHTTP(recorder, request)

// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func TestCarthandler_GetCart(t *testing.T) {

// }

// func TestCarthandler_DeleteCartItem(t *testing.T) {

// }

// func TestCarthandler_DeleteCartAll(t *testing.T) {

// }
