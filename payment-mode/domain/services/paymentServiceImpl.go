package services

import (
	"net/http"
	app_errors "qwik.in/payment-mode/app-errors"
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

func (p PaymentServiceImpl) GetUserId(token string) string {

	//TODO
	//Should communicate with the authentication service to fetch userId from the token passed.
	return token
}

func (p PaymentServiceImpl) AddPaymentMode(paymentMode *models.PaymentMode, userId string) *app_errors.AppError {

	//Fetch userPaymentMode record from DB.
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(userId)

	//If there is no record for the given userId, create a new record.
	if err != nil {
		if err.Code == http.StatusNotFound {
			newPaymentModes := make([]models.PaymentMode, 0, 0)
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

func (p PaymentServiceImpl) SetPaymentMode(userId string, paymentMode models.PaymentMode) (bool, *app_errors.AppError) {
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(userId)

	if err != nil {
		return false, err
	}

	for _, addedPaymentModes := range userPaymentModes.PaymentModes {
		if addedPaymentModes.Mode == paymentMode.Mode && addedPaymentModes.CardNumber == paymentMode.CardNumber {
			return true, nil
		}
	}

	return false, app_errors.NewNotFoundError("The given payment mode doesn't exist for the current user.")

}

func (p PaymentServiceImpl) CheckBalanceAndCompletePayment(paymentRequest *models.PaymentRequest) (bool, *app_errors.AppError) {
	userPaymentModes, err := p.paymentRepository.GetPaymentModeFromDB(paymentRequest.UserId)

	if err != nil {
		return false, err
	}

	for i, availablePaymentMode := range userPaymentModes.PaymentModes {

		if availablePaymentMode.Mode == paymentRequest.SelectedPaymentMode.Mode && availablePaymentMode.CardNumber == paymentRequest.SelectedPaymentMode.CardNumber {

			if availablePaymentMode.Balance >= paymentRequest.OrderAmount {

				availablePaymentMode.Balance -= paymentRequest.OrderAmount
				userPaymentModes.PaymentModes[i] = availablePaymentMode
				updateErr := p.paymentRepository.UpdatePaymentModeToDB(userPaymentModes)

				if updateErr != nil {
					return false, updateErr
				} else {
					return true, nil
				}

			} else {
				return false, app_errors.NewRequestNotAcceptedError("Insufficient funds, payment failed.")
			}
		}
	}

	return false, app_errors.NewNotFoundError("Payment method is not added for the current user.")
}
