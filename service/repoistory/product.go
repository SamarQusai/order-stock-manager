package repoistory

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) GetProductByObjectsIds(sessionContext mongo.SessionContext, productsIds []primitive.ObjectID) ([]db_model.Product, error) {
	var productsObject []db_model.Product
	cursor, findProductsError := c.db.Collection(db_model.ProductCollectionName).Find(sessionContext, bson.M{"_id": bson.M{"$in": productsIds}})
	if findProductsError != nil {
		c.context.Logger().Error("Error while finding products, error: ", findProductsError.Error())
	}
	decodingError := cursor.All(context.Background(), &productsObject)
	if decodingError != nil {
		c.context.Logger().Error("Decoding error while finding products - decodingErr, ", decodingError.Error())
		return nil, decodingError
	}
	return productsObject, nil
}
