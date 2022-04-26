package models

type UserPaymentMode struct {
	UserId       string        `json:"user_id" dynamodbav:"user_id,omitempty"`
	PaymentModes []PaymentMode `json:"payment_modes" dynamodbav:"payment_modes" validate:"required"`
}
