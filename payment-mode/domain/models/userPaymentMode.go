package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPaymentMode struct {
	UserId       primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	PaymentModes []PaymentMode      `json:"payment_modes" bson:"payment_modes" validate:"required"`
}
