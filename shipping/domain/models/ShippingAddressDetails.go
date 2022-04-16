package models

//ShippingAddressDetails ..
type ShippingAddressDetails struct {
	ShippingAddressID int    `json:"shippingAddressId" dynamodb:"shippingAddressId" validate:"required"`
	FirstName         string `json:"firstName" dynamodbav:"firstName" validate:"required"`
	LastName          string `json:"lastName" dynamodbav:"lastName"`
	AddressLine1      string `json:"addressLine1" dynamodbav:"addressLine1" validate:"required"`
	AddressLine2      string `json:"addressLine2" dynamodbav:"addressLine2" validate:"required"`
	City              string `json:"city" dynamodbav:"city" validate:"required"`
	State             string `json:"state" dynamodbav:"state" validate:"required"`
	Phone             string `json:"phone" dynamodbav:"phone" validate:"required"`
	Pincode           string `json:"pincode" dynamodbav:"pincode" validate:"required"`
	AddressType       string `json:"addressType,omitempty" dynamodbav:"addressType" validate:"required"`
	DefaultAddress    bool   `json:"defaultAddress,omitempty" dynamodbav:"defaultAddress"`
}
