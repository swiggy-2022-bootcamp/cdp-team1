package models

type PaymentRequest struct {
	UserId              string      `json:"user_id" validate:"required"`
	OrderId             string      `json:"order_id" validate:"required"`
	OrderAmount         int         `json:"order_amount" validate:"required"`
	SelectedPaymentMode PaymentMode `json:"payment_mode" validate:"required"`
}
