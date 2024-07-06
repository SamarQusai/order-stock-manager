package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) PlaceOrderProcessing(request model.PlaceOrderProcessingRequest) (string, error) {

	session, startingSessionError := c.db.Client().StartSession()
	if startingSessionError != nil {
		log.Fatal(startingSessionError)
	}
	defer session.EndSession(context.Background())
	transactionOptions := options.Transaction()
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		products, getProductDetailsError := c.GetProductDetails(sessionContext, request.ProductIds)
		if getProductDetailsError != nil {
			return nil, getProductDetailsError
		}
		updatedIngredientsStock, ingredientsAvailabilityError := c.CheckIngredientsAvailability(sessionContext, products, request)
		if ingredientsAvailabilityError != nil {
			return nil, ingredientsAvailabilityError
		}

		for _, ingredient := range updatedIngredientsStock {
			decreasingError := c.DecreaseStock(sessionContext, ingredient)
			if decreasingError != nil {
				return nil, decreasingError
			}
		}

		var orderItems []db_model.Item
		for _, productToAdd := range request.Products {
			orderItems = append(orderItems, db_model.Item{
				ID:    productToAdd.ProductId,
				Price: 0,
			})
		}
		var order = db_model.Order{
			ID:    primitive.ObjectID{},
			Items: orderItems,
		}
		id, placingOrderError := c.PlaceOrder(sessionContext, order)
		if placingOrderError != nil {
			return nil, placingOrderError
		}

		return model.PlaceOrderProcessingResponse{
			OrderId:     id.Hex(),
			Ingredients: updatedIngredientsStock,
		}, nil
	}

	result, err := session.WithTransaction(context.Background(), callback, transactionOptions)
	if err != nil {
		c.context.Logger().Error("Failed to process creating order transaction, error: ", err.Error())
		return "", err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				c.context.Logger().Error(`error while sending email, error: `, r)
			}
		}()
		c.notifyMerchant(result.(model.PlaceOrderProcessingResponse).Ingredients)
		return
	}()

	return result.(model.PlaceOrderProcessingResponse).OrderId, nil
}

func (c *Config) PlaceOrder(sessionContext mongo.SessionContext, order db_model.Order) (primitive.ObjectID, error) {
	id, err := c.repository.PersistOrder(sessionContext, order)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return id, nil
}
