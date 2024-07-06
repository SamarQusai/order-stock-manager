package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/context"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/model/db_model"
	"simple-order-stock-manager/service/repoistory"
)

type Interface interface {
	GetProductDetails(sessionContext mongo.SessionContext, products []primitive.ObjectID) ([]db_model.Product, error)
	CheckIngredientsAvailability(sessionContext mongo.SessionContext, products []db_model.Product, request model.PlaceOrderProcessingRequest) ([]db_model.Ingredient, error)
	DecreaseStock(sessionContext mongo.SessionContext, ingredient db_model.Ingredient) error
	PlaceOrderProcessing(request model.PlaceOrderProcessingRequest) (string, error)
	PlaceOrder(sessionContext mongo.SessionContext, order db_model.Order) (primitive.ObjectID, error)
}

type Config struct {
	context    context.ServiceContext
	repository repoistory.Interface
	db         *mongo.Database
}

func New(context context.ServiceContext) *Config {
	return &Config{
		context:    context,
		repository: repoistory.New(context),
		db:         context.GetDB(),
	}
}
