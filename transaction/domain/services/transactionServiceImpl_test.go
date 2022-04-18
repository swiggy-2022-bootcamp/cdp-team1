package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	app_erros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/mocks"
	"reflect"
	"testing"
)

func TestTransactionServiceImpl_AddTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transactionAmount := &models.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		OrderId: "OA-123",
		Amount:  500,
	}
	transaction := &models.Transaction{
		TransactionPoints: 5,
		UserId:            transactionAmount.UserId,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, expected interface{}, actual interface{})
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
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
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
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
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
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), actual)
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
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), actual)
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
			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), actual)
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

			transactionServiceImpl := NewTransactionServiceImpl(transactionRepository)
			err := transactionServiceImpl.AddTransactionPoints(transactionAmount)
			tc.checkResponse(t, nil, err)
		})
	}
}

func TestTransactionServiceImpl_GetTransactionPointsByUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	userId := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	failurePoints := -1
	successPoints := 100

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, points interface{}, err interface{})
	}{
		{
			name: "SuccessUserFound",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(userId).
					Times(1).
					Return(successPoints, nil)
			},
			checkResponse: func(t *testing.T, points interface{}, err interface{}) {
				require.Equal(t, successPoints, points)
				require.Equal(t, true, reflect.ValueOf(err).IsNil())
			},
		},
		{
			name: "FailureUserNotFound",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(userId).
					Times(1).
					Return(failurePoints, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, points interface{}, err interface{}) {
				require.Equal(t, app_erros.NewNotFoundError(""), err)
				require.Equal(t, failurePoints, points)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(userId).
					Times(1).
					Return(failurePoints, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, points interface{}, err interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), err)
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

			transactionServiceImpl := NewTransactionServiceImpl(transactionRepository)
			points, err := transactionServiceImpl.GetTransactionPointsByUserId(userId)
			tc.checkResponse(t, points, err)
		})
	}
}

func TestTransactionServiceImpl_CalculateTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testCases := []struct {
		name              string
		transactionAmount *models.TransactionDetails
		expectedPoints    int
	}{
		{
			name: "Whole number",
			transactionAmount: &models.TransactionDetails{
				UserId: "1",
				Amount: 5000,
			},
			expectedPoints: 50,
		},
		{
			name: "Amount less than 100",
			transactionAmount: &models.TransactionDetails{
				UserId: "1",
				Amount: 5,
			},
			expectedPoints: 0,
		},
		{
			name: "Points rounded to narest integer",
			transactionAmount: &models.TransactionDetails{
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
			transactionService := NewTransactionServiceImpl(repository)
			actual := transactionService.CalculateTransactionPoints(tc.transactionAmount)
			require.Equal(t, tc.expectedPoints, actual)
		})
	}
}

func TestTransactionServiceImpl_UseTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transactionAmount := &models.TransactionDetails{
		UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		Amount:  5000,
		OrderId: "OA-123",
	}

	transaction := &models.Transaction{
		UserId:            transactionAmount.UserId,
		TransactionPoints: 0,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{})
	}{
		{
			name: "FailureGetTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionAmount.UserId).
					Times(1).
					Return(-1, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{}) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionAmount)
				require.Equal(t, app_erros.NewUnexpectedError(""), err)
			},
		},
		{
			name: "UserTransactionPointsIsZero",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionAmount.UserId).
					Times(1).
					Return(0, nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{}) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionAmount)
				require.Equal(t, app_erros.NewExpectationFailed("You have 0 transaction points"), err)
			},
		},
		{
			name: "Success",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionAmount.UserId).
					Times(1).
					Return(100, nil)

				transactionAmount.Amount -= 100

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{}) {
				require.Equal(t, true, isRequestValid)
				require.Equal(t, amount, transactionAmount)
				require.Equal(t, true, reflect.ValueOf(err).IsNil())
			},
		},
		{
			name: "FailureUpdateTransactionToDB",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionAmount.UserId).
					Times(1).
					Return(100, nil)

				transactionAmount.Amount -= 100

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{}) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionAmount)
				require.Equal(t, app_erros.NewUnexpectedError(""), err)
			},
		},
		{
			name: "OrderAmountLessThanTransactionPoints",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					GetTransactionPointsByUserIdFromDB(transactionAmount.UserId).
					Times(1).
					Return(100000, nil)

				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, isRequestValid bool, amount *models.TransactionDetails, err interface{}) {
				require.Equal(t, false, isRequestValid)
				require.Equal(t, amount, transactionAmount)
				require.Equal(t, app_erros.NewExpectationFailed("Cannot use transaction points as order amount is lesser than available points"), err)
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

			transactionServiceImpl := NewTransactionServiceImpl(repository)

			requestAccepted, amount, err := transactionServiceImpl.UseTransactionPoints(transactionAmount)
			tc.checkResponse(t, requestAccepted, amount, err)
		})
	}
}

func TestTransactionServiceImpl_UpdateTransactionPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	transaction := &models.Transaction{
		UserId:            "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		TransactionPoints: 100,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockTransactionRepository)
		checkResponse func(t *testing.T, actual interface{}, expected interface{})
	}{
		{
			name: "Success",
			buildStubs: func(repository *mocks.MockTransactionRepository) {
				repository.EXPECT().
					UpdateTransactionPointsToDB(transaction).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, true, reflect.ValueOf(expected).IsNil())
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
			checkResponse: func(t *testing.T, actual interface{}, expected interface{}) {
				require.Equal(t, app_erros.NewUnexpectedError(""), expected)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repostiory := mocks.NewMockTransactionRepository(ctrl)
			transactionServiceImpl := NewTransactionServiceImpl(repostiory)

			tc.buildStubs(repostiory)

			err := transactionServiceImpl.UpdateTransactionPoints(transaction.TransactionPoints, transaction.UserId)
			tc.checkResponse(t, nil, err)
		})
	}
}
