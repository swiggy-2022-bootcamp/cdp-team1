package models

//PaymentUserPaymentModeModel - Payment Service
type PaymentUserPaymentModeModel struct {
	OrderAmount      float64 `json:"order_amount" dynamodbav:"order_amount,omitempty"`
	UserID           string  `json:"user_id" dynamodbav:"user_id,omitempty"`
	PaymentModeModel []struct {
		Mode       string `json:"mode_of_payment" dynamodbav:"mode_of_payment"`
		Credential string `json:"credential" dynamodbav:"credential"`
	} `json:"payment_mode" dynamodbav:"payment_mode"`
}

/*
Create a JSON Body - POST the payment details -> card number,

1. Random OrderID ->
*/
