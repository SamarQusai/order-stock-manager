package repoistory

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) PersistOrder(sessionContext mongo.SessionContext, order db_model.Order) (primitive.ObjectID, error) {
	result, insertingError := c.db.Collection(db_model.OrderCollectionName).InsertOne(sessionContext, order)
	if insertingError != nil {
		c.context.Logger().Error("Error while persisting order, error: ", insertingError.Error())
		return primitive.NilObjectID, insertingError
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
