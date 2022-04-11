package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/log"
)

type PaymentRepositoryImpl struct {
	mongoCollection *mongo.Collection
	ctx             context.Context
}

func NewPaymentRepositoryImpl(mongoCollection *mongo.Collection, ctx context.Context) PaymentRepository {
	return &PaymentRepositoryImpl{
		mongoCollection: mongoCollection,
		ctx:             ctx,
	}
}

func (p PaymentRepositoryImpl) AddPaymentModeToDB(userPaymentMode *models.UserPaymentMode) error {
	_, err := p.mongoCollection.InsertOne(p.ctx, userPaymentMode)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (p PaymentRepositoryImpl) GetPaymentModeFromDB(userId string) (*models.UserPaymentMode, error) {
	var userPaymentMode models.UserPaymentMode
	objId, _ := primitive.ObjectIDFromHex(userId)
	query := bson.D{bson.E{Key: "user_id", Value: objId}}
	err := p.mongoCollection.FindOne(p.ctx, query).Decode(&userPaymentMode)
	return &userPaymentMode, err
}
