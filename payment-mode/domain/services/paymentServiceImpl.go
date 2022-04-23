package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/config"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/domain/repository"
	"qwik.in/payment-mode/log"
)

type PaymentServiceImpl struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentServiceImpl(paymentRepository repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepository: paymentRepository,
	}
}

func (p PaymentServiceImpl) GetUserId(authToken string) (string, *app_errors.AppError) {

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.EnvJWTSecretKey()), nil
	})
	if err != nil {
		return "", app_errors.NewAuthenticationError("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims["user_id"].(string), nil
	}
	return "", app_errors.NewAuthenticationError("invalid token")
}

func (p PaymentServiceImpl) AddPaymentMode(paymentMode *models.PaymentMode, userId string) *app_errors.AppError {

	//Fetch userPaymentMode record from DB.
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(userId)

	//If there is no record for the given userId, create a new record.
	if err != nil {
		if err.Code == http.StatusNotFound {
			newPaymentModes := make([]models.PaymentMode, 0, 0)
			//Generating random number to initialize balance
			paymentMode.Balance = 50000
			newUserPaymentMode := models.UserPaymentMode{
				UserId:       userId,
				PaymentModes: append(newPaymentModes, *paymentMode),
			}
			addPaymentErr := p.paymentRepository.AddPaymentModeToDB(&newUserPaymentMode)
			if addPaymentErr != nil {
				return addPaymentErr
			} else {
				return nil
			}
		} else {
			return err
		}
	}

	//Record with userId already exists, so adding the payment mode to the user.
	//Check if passed paymentMode object is there or not in userPaymentMode, if there return Conflict
	currentPaymentModes := userPaymentModes.PaymentModes
	for _, currentPaymentMode := range currentPaymentModes {
		if currentPaymentMode.Mode == paymentMode.Mode && currentPaymentMode.CardNumber == paymentMode.CardNumber {
			newErr := app_errors.NewConflictRequestError("This payment mode is already present in your account.")
			return newErr
		}
	}

	// If not there add the payment mode object in userPaymentMode's paymentMethods field.
	paymentMode.Balance = 50000
	currentPaymentModes = append(currentPaymentModes, *paymentMode)
	newUserPaymentMode := models.UserPaymentMode{
		UserId:       userId,
		PaymentModes: currentPaymentModes,
	}
	err = p.paymentRepository.UpdatePaymentModeToDB(&newUserPaymentMode)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p PaymentServiceImpl) GetPaymentMode(userId string) (*models.UserPaymentMode, *app_errors.AppError) {

	// Fetch payment modes for the given user
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return userPaymentModes, nil
}

func (p PaymentServiceImpl) CheckBalanceAndCompletePayment(paymentRequest *models.PaymentRequest) (bool, *app_errors.AppError) {
	// Fetch available modes of payment for the given user.
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(paymentRequest.UserId)

	if err != nil {
		return false, err
	}

	// Looping through the available payment modes
	for i, availablePaymentMode := range userPaymentModes.PaymentModes {

		// Searching for selected payment mode from available modes
		if availablePaymentMode.Mode == paymentRequest.SelectedPaymentMode.Mode && availablePaymentMode.CardNumber == paymentRequest.SelectedPaymentMode.CardNumber {

			// Verifying if the available balance is sufficient for the given order amount.
			if availablePaymentMode.Balance >= paymentRequest.OrderAmount {

				// Reducing the balance and updating it to the DB.
				availablePaymentMode.Balance -= paymentRequest.OrderAmount
				userPaymentModes.PaymentModes[i] = availablePaymentMode
				updateErr := p.paymentRepository.UpdatePaymentModeToDB(userPaymentModes)

				// Payment failure. Returning false due to DB update failure
				if updateErr != nil {
					return false, updateErr
				} else {
					return true, nil
				}

			} else {
				//  Payment failure. Balance is less than the order amount.
				return false, app_errors.NewRequestNotAcceptedError("Insufficient funds, payment failed.")
			}
		}
	}

	// Payment failure. Payment mode not found in user payment mode list.
	return false, app_errors.NewNotFoundError("Payment method is not added for the current user.")
}
