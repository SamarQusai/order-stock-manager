package repoistory

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/context"
	"simple-order-stock-manager/model/db_model"
)

type Interface interface {
	GetProductByObjectsIds(sessionContext mongo.SessionContext, products []primitive.ObjectID) ([]db_model.Product, error)
	GetIngredientById(sessionContext mongo.SessionContext, id primitive.ObjectID) (*db_model.Ingredient, error)
	GetIngredientsByIds(sessionContext mongo.SessionContext, ids []primitive.ObjectID) ([]db_model.Ingredient, error)
	Decrease(sessionContext mongo.SessionContext, id primitive.ObjectID, stock float64) error
	PersistOrder(sessionContext mongo.SessionContext, order db_model.Order) (primitive.ObjectID, error)
	FindEmailByResourceId(resourceId primitive.ObjectID) (*db_model.SentEmail, error)
	PersistEmail(email db_model.SentEmail) error
}

type Config struct {
	context context.ServiceContext
	db      *mongo.Database
}

func New(context context.ServiceContext) *Config {
	return &Config{
		context: context,
		db:      context.GetDB(),
	}
}
