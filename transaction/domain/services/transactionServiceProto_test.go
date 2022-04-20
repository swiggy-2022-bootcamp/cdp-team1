package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	app_erros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/mocks"
	"qwik.in/transaction/protos"
	"reflect"
	"testing"
)

func TestTransactionProtoServer_AddTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transactionDetails := &protos.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		OrderId: "OA-123",
		Amount:  500,
	}
	transaction := &models.Transaction{
		TransactionPoints: 5,
		UserId:            transactionDetails.UserId,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, err error)
	}{
		{
			name: "SuccessUpdateTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transaction.UserId).
					Times(1).
					Return(10, nil)
				transaction.TransactionPoints = 15
				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "SuccessAddTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transaction.UserId).
					Times(1).
					Return(-1, app_erros.NewNotFoundError(""))
				transaction.TransactionPoints = 5
				repository.EXPECT().
					AddTransactionPointsFromDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "FailureAddTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transaction.UserId).
					Times(1).
					Return(-1, app_erros.NewNotFoundError(""))

				repository.EXPECT().
					AddTransactionPointsFromDB(transaction).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
			},
		},
		{
			name: "FailureGetTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transaction.UserId).
					Times(1).
					Return(-1, app_erros.NewUnexpectedError(""))

				repository.EXPECT().
					AddTransactionPointsFromDB(transaction).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
			},
		},
		{
			name: "FailureUpdateTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transaction.UserId).
					Times(1).
					Return(10, nil)
				transaction.TransactionPoints = 15
				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			transactionRepository := mocks.NewMockTransactionRepository(ctrl)
			tc.buildStubs(transactionRepository)

			transactionProtoServer := NewTransactionProtoServer(transactionRepository)
			_, err := transactionProtoServer.AddTransactionPoints(context.Background(), transactionDetails)
			tc.checkResponse(t, err)
		})
	}
}

func TestTransactionProtoServer_GetTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := &protos.GetPointsRequest{
		UserId: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	}
	failurePoints := &protos.GetPointsResponse{
		Points: int32(-1),
	}
	successPoints := &protos.GetPointsResponse{
		Points: int32(100),
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, points *protos.GetPointsResponse, err error)
	}{
		{
			name: "SuccessUserFound",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(request.GetUserId()).
					Times(1).
					Return(100, nil)
			},
			checkResponse: func(t *testing.T, points *protos.GetPointsResponse, err error) {
				require.Equal(t, successPoints, points)
				require.NoError(t, err)
			},
		},
		{
			name: "FailureUserNotFound",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(request.GetUserId()).
					Times(1).
					Return(-1, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, points *protos.GetPointsResponse, err error) {
				require.Equal(t, app_erros.NewNotFoundError("").Error(), err)
				require.Equal(t, failurePoints, points)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(request.GetUserId()).
					Times(1).
					Return(-1, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, points *protos.GetPointsResponse, err error) {
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
				require.Equal(t, failurePoints, points)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			transactionRepository := mocks.NewMockTransactionRepository(ctrl)
			tc.buildStubs(transactionRepository)

			transactionProtoServer := NewTransactionProtoServer(transactionRepository)
			points, err := transactionProtoServer.GetTransactionPoints(context.Background(), request)
			tc.checkResponse(t, points, err)
		})
	}
}

func TestTransactionProtoServer_UseTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transactionDetails := &protos.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		Amount:  5000,
		OrderId: "OA-123",
	}

	transaction := &models.Transaction{
		UserId:            transactionDetails.UserId,
		TransactionPoints: 0,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error)
	}{
		{
			name: "FailureGetTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionDetails.UserId).
					Times(1).
					Return(-1, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionDetails)
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
			},
		},
		{
			name: "UserTransactionPointsIsZero",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionDetails.UserId).
					Times(1).
					Return(0, nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionDetails)
				require.Equal(t, app_erros.NewExpectationFailed("You have 0 transaction points").Error(), err)
			},
		},
		{
			name: "Success",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionDetails.UserId).
					Times(1).
					Return(100, nil)

				transactionDetails.Amount -= 100

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error) {
				require.Equal(t, true, isRequestValid)
				require.Equal(t, amount, transactionDetails)
				require.NoError(t, err)
			},
		},
		{
			name: "FailureUpdateTransactionToDB",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionDetails.UserId).
					Times(1).
					Return(100, nil)

				transactionDetails.Amount -= 100

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionDetails)
				require.Equal(t, app_erros.NewUnexpectedError("").Error(), err)
			},
		},
		{
			name: "OrderAmountLessThanTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionDetails.UserId).
					Times(1).
					Return(100000, nil)

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *protos.TransactionDetails, err error) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionDetails)
				require.Equal(t, app_erros.NewExpectationFailed("Cannot use transaction points as order amount is lesser than available points").Error(), err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockTransactionRepository(ctrl)
			tc.buildStubs(repository)

			transactionProtoServer := NewTransactionProtoServer(repository)

			response, err := transactionProtoServer.UseTransactionPoints(context.Background(), transactionDetails)
			tc.checkResponse(t, response.TransactionPointsUsed, response.TransactionDetails, err)
		})
	}
}

func TestTransactionProtoServer_CalculateTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testCases := []struct {
		name              string
		transactionAmount *protos.TransactionDetails
		expectedPoints    int
	}{
		{
			name: "Whole number",
			transactionAmount: &protos.TransactionDetails{
				UserId: "1",
				Amount: 5000,
			},
			expectedPoints: 50,
		},
		{
			name: "Amount less than 100",
			transactionAmount: &protos.TransactionDetails{
				UserId: "1",
				Amount: 5,
			},
			expectedPoints: 0,
		},
		{
			name: "Points rounded to narest integer",
			transactionAmount: &protos.TransactionDetails{
				UserId: "1",
				Amount: 399,
			},
			expectedPoints: 3,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockTransactionRepository(ctrl)
			transactionProtoServer := NewTransactionProtoServer(repository)
			actual := transactionProtoServer.CalculateTransactionPoints(tc.transactionAmount)
			require.Equal(t, tc.expectedPoints, actual)
		})
	}
}

func TestTransactionProtoServer_UpdateTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transaction := &models.Transaction{
		UserId:            "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		TransactionPoints: 100,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, err interface{})
	}{
		{
			name: "Success",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, err interface{}) {
				require.Equal(t, true, reflect.ValueOf(err).IsNil())
			},
		},
		{
			name: "Failure",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, err interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repostiory := mocks.NewMockTransactionRepository(ctrl)
			transactionProtoServer := NewTransactionProtoServer(repostiory)

			tc.buildStubs(repostiory)

			err := transactionProtoServer.UpdateTransactionPoints(transaction.TransactionPoints, transaction.UserId)
			tc.checkResponse(t, err)
		})
	}
}
