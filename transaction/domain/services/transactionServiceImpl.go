package services

import (
	"net/http"
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/domain/repository"
	"qwik.in/transaction/log"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepository: transactionRepository}
}

func (t TransactionServiceImpl) AddTransactionPoints(transactionAmount *models.TransactionDetails) *apperros.AppError {
	//calculate transaction points earned
	points := t.CalculateTransactionPoints(transactionAmount)
	transaction := &models.Transaction{
		TransactionPoints: points,
		UserId:            transactionAmount.UserId,
	}

	//Fetch user transaction details from DB
	userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(transaction.UserId)

	//If there is no record for the given userId, create a new record
	if err != nil {
		if err.Code == http.StatusNotFound {
			err_ := t.transactionRepository.AddTransactionPointsFromDB(transaction)
			if err_ != nil {
				return err_
			} else {
				return nil
			}
		} else {
			return err
		}
	}

	//Record with userId already exists, so updating the transaction points
	transaction.TransactionPoints += userTransactionPoints
	err = t.transactionRepository.UpdateTransactionPointsToDB(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionServiceImpl) GetTransactionPointsByUserId(userId string) (int, *apperros.AppError) {
	userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(userId)
	if err != nil {
		return -1, err
	}
	return userTransactionPoints, nil
}

func (t TransactionServiceImpl) CalculateTransactionPoints(transactionAmount *models.TransactionDetails) int {
	//For every 100 Rupees user will get 1 transaction point
	var points int
	points = transactionAmount.Amount / 100
	return points
}

func (t TransactionServiceImpl) UseTransactionPoints(transactionAmount *models.TransactionDetails) (bool, *models.TransactionDetails, *apperros.AppError) {
	//Fetch user transaction details from DB
	userTransactionPoints, err := t.transactionRepository.GetTransactionPointsByUserIdFromDB(transactionAmount.UserId)

	if err != nil {
		return false, transactionAmount, err
	} else {
		// For every 1 transaction point, user will get a discount of 1 Rupee
		if userTransactionPoints == 0 {
			return false, transactionAmount, apperros.NewExpectationFailed("You have 0 transaction points")
		} else if userTransactionPoints < transactionAmount.Amount {
			//user using all his transaction points (Will only apply if order amount is greater than points)
			transactionAmount.Amount -= userTransactionPoints
			userTransactionPoints = 0

			err_ := t.UpdateTransactionPoints(userTransactionPoints, transactionAmount.UserId)
			if err_ != nil {
				log.Error("Failed to update transaction points")
				return false, transactionAmount, err_
			}

			return true, transactionAmount, nil
		} else {
			log.Info("Cannot use transaction points as order amount is lesser than available points")
			return false, transactionAmount, apperros.NewExpectationFailed("Cannot use transaction points as order amount is lesser than available points")
		}
	}
}

func (t TransactionServiceImpl) UpdateTransactionPoints(transactionPoint int, userId string) *apperros.AppError {
	transaction := &models.Transaction{
		UserId:            userId,
		TransactionPoints: transactionPoint,
	}
	err := t.transactionRepository.UpdateTransactionPointsToDB(transaction)
	return err
}
