package models

//ShippingAddrModel ..
type ShippingAddrModel struct {
	UserID            int    `json:"user_id" dynamodbav:"user_id"`
	ShippingAddressID string `json:"shipping_address_id" dynamodbav:"shipping_address_id"`
	FirstName         string `json:"first_name" dynamodbav:"first_name"`
	LastName          string `json:"last_name" dynamodbav:"lastName"`
	AddressLine1      string `json:"address_line_1" dynamodbav:"address_line_1" `
	AddressLine2      string `json:"address_line_2" dynamodbav:"address_line_2"`
	City              string `json:"city" dynamodbav:"city"`
	State             string `json:"state" dynamodbav:"state"`
	Phone             string `json:"phone" dynamodbav:"phone"`
	Pincode           int    `json:"pincode" dynamodbav:"pincode"`
	AddressType       string `json:"address_type" dynamodbav:"address_type"`
	DefaultAddress    bool   `json:"default_address" dynamodbav:"default_address"`
	ShippingCost      int    `json:"shipping_cost" dynamodb:"shipping_cost"`
}
