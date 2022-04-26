package services

import (
	"context"
	"net/http"
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/domain/repository"
	"qwik.in/transaction/log"
	"qwik.in/transaction/protos"
)

var transactionRepository repository.TransactionRepository

type TransactionProtoServer struct {
	protos.UnimplementedTransactionPointsServer
}

func NewTransactionProtoServer(tr repository.TransactionRepository) TransactionProtoServer {
	transactionRepository = tr
	return TransactionProtoServer{}
}

func (s TransactionProtoServer) AddTransactionPoints(ctx context.Context, transactionDetails *protos.TransactionDetails) (*protos.AddPointsResponse, error) {
	addPointsResponse := &protos.AddPointsResponse{}
	//calculate transaction points earned
	points := s.CalculateTransactionPoints(transactionDetails)
	transaction := &models.Transaction{
		TransactionPoints: points,
		UserId:            transactionDetails.GetUserId(),
	}

	//Fetch user transaction details from DB
	userTransactionPoints, err := transactionRepository.GetTransactionPointsByUserIdFromDB(transaction.UserId)

	//If there is no record for the given userId, create a new record
	if err != nil {
		if err.Code == http.StatusNotFound {
			dbErr := transactionRepository.AddTransactionPointsFromDB(transaction)
			if dbErr != nil {
				return addPointsResponse, dbErr.Error()
			} else {
				return addPointsResponse, nil
			}
		} else {
			return addPointsResponse, err.Error()
		}
	}

	//Record with userId already exists, so updating the transaction points
	transaction.TransactionPoints += userTransactionPoints
	err = transactionRepository.UpdateTransactionPointsToDB(transaction)
	if err != nil {
		return addPointsResponse, err.Error()
	}
	return addPointsResponse, nil
}

func (s TransactionProtoServer) UseTransactionPoints(ctx context.Context, transactionDetails *protos.TransactionDetails) (*protos.UsePointsResponse, error) {
	usePointsResponse := &protos.UsePointsResponse{
		TransactionPointsUsed: false,
		TransactionDetails:    transactionDetails,
	}
	userTransactionPoints, err := transactionRepository.GetTransactionPointsByUserIdFromDB(transactionDetails.GetUserId())

	if err != nil {
		return usePointsResponse, err.Error()
	} else {
		// For every 1 transaction point, user will get a discount of 1 Rupee
		if userTransactionPoints == 0 {
			return usePointsResponse, apperros.NewExpectationFailed("You have 0 transaction points").Error()
		} else if userTransactionPoints < int(transactionDetails.Amount) {
			//user using all his transaction points (Will only apply if order amount is greater than points)
			transactionDetails.Amount -= int32(userTransactionPoints)
			userTransactionPoints = 0

			updateErr := s.UpdateTransactionPoints(userTransactionPoints, transactionDetails.GetUserId())
			if updateErr != nil {
				log.Error("Failed to update transaction points")
				return usePointsResponse, updateErr.Error()
			}
			usePointsResponse.TransactionPointsUsed = true
			usePointsResponse.TransactionDetails = transactionDetails
			return usePointsResponse, nil
		} else {
			log.Info("Cannot use transaction points as order amount is lesser than available points")
			return usePointsResponse, apperros.NewExpectationFailed("Cannot use transaction points as order amount is lesser than available points").Error()
		}
	}
}

func (s TransactionProtoServer) GetTransactionPoints(ctx context.Context, request *protos.GetPointsRequest) (*protos.GetPointsResponse, error) {
	userTransactionPoints, err := transactionRepository.GetTransactionPointsByUserIdFromDB(request.GetUserId())
	if err != nil {
		return &protos.GetPointsResponse{Points: int32(-1)}, err.Error()
	}
	return &protos.GetPointsResponse{Points: int32(userTransactionPoints)}, nil
}

func (s TransactionProtoServer) CalculateTransactionPoints(transactionDetails *protos.TransactionDetails) int {
	//For every 100 Rupees user will get 1 transaction point
	var points int
	points = int(transactionDetails.GetAmount() / 100)
	return points
}

func (s TransactionProtoServer) UpdateTransactionPoints(transactionPoint int, userId string) *apperros.AppError {
	transaction := &models.Transaction{
		UserId:            userId,
		TransactionPoints: transactionPoint,
	}
	err := transactionRepository.UpdateTransactionPointsToDB(transaction)
	return err
}
