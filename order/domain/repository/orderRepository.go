package repository

import (
	"orderService/domain/model"
	"orderService/internal/error"
)

type OrderRepositoryDB interface {
	ReadStatus(string) (*[]model.Order, *error.AppError)
	ReadID(string) (*model.Order, *error.AppError)
	ReadCustomerID(string) (*model.Order, *error.AppError)
	ReadAll() (*[]model.Order, *error.AppError)
	Update(model.Order) (*model.Order, *error.AppError)
	Delete(model.Order) (*model.Order, *error.AppError)
	DeleteAll() *error.AppError
}

type OrderRepository struct {
	// dbClient *mongo.Client
}

func (odb OrderRepository) ReadStatus(id string) (*[]model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &[]model.Order{}, nil
}

func (odb OrderRepository) ReadID(id string) (*model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Order{}, nil
}

func (odb OrderRepository) ReadCustomerID(id string) (*model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Order{}, nil
}

func (odb OrderRepository) ReadAll() (*[]model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &[]model.Order{}, nil
}

func (odb OrderRepository) Update(order model.Order) (*model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Order{}, nil
}

func (odb OrderRepository) Delete(order model.Order) (*model.Order, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "Order")
	// order := model.Order{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Order{}, nil
}
