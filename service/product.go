package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) GetProductDetails(sessionContext mongo.SessionContext, productsIds []primitive.ObjectID) ([]db_model.Product, error) {
	products, getProductsError := c.repository.GetProductByObjectsIds(sessionContext, productsIds)

	if getProductsError != nil {
		c.context.Logger().Error("Error while finding products, error: ", getProductsError.Error())
		if getProductsError.Error() == mongo.ErrNoDocuments.Error() {
			return nil, model.ProductNotFound
		}
		return nil, getProductsError
	}

	if len(products) != len(productsIds) {
		c.context.Logger().Error("Products length don't match the products in the request")
		return nil, model.ProductNotFound
	}

	return products, nil
}
