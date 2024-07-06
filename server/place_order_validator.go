package server

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"simple-order-stock-manager/model"
)

func (c *Config) validate(httpRequest model.PlaceOrderHTTPRequest) (*model.PlaceOrderProcessingRequest, error) {
	var internalRequest model.PlaceOrderProcessingRequest
	var productIds []primitive.ObjectID
	var productDetails []model.ProductRequest
	productsQuantities := make(map[primitive.ObjectID]int32)
	for _, product := range httpRequest.Products {
		productObjectId, conversionError := primitive.ObjectIDFromHex(product.ProductId)
		if conversionError != nil {
			c.context.Logger().Error("Error while converting product id to object id, id: ", product.ProductId, " error: ", conversionError.Error())
			return nil, errors.New("invalid product id")
		}
		productsQuantities[productObjectId] += product.Quantity
		productIds = append(productIds, productObjectId)
		productDetails = append(productDetails, model.ProductRequest{
			ProductId: productObjectId,
			Quantity:  product.Quantity,
		})
	}
	internalRequest.ProductIds = productIds
	internalRequest.ProductsQuantities = productsQuantities
	internalRequest.Products = productDetails
	return &internalRequest, nil
}
