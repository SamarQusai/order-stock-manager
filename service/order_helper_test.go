package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/utils"
	"testing"
)

func Test_PlaceOrderProcessingSuccessfullyCreating(t *testing.T) {
	productObjectId, _ := primitive.ObjectIDFromHex("6688612be29cfdf30d5db4be")
	productIds := make([]primitive.ObjectID, 0)
	productIds = append(productIds, productObjectId)
	var products []model.ProductRequest
	productsQuantities := make(map[primitive.ObjectID]int32, 0)
	productsQuantities[productObjectId] = 1
	products = append(products, model.ProductRequest{
		ProductId: productObjectId,
		Quantity:  1,
	})
	id, err := services.PlaceOrderProcessing(model.PlaceOrderProcessingRequest{
		ProductIds:         productIds,
		ProductsQuantities: productsQuantities,
		Products:           products,
	})

	if err != nil {
		t.Fatal(`Error while placing order`, err)
	}
	fmt.Println("order id", id)
}

func Test_PlaceOrderProcessingOutOfStock(t *testing.T) {
	productObjectId, _ := primitive.ObjectIDFromHex("66891e9de29cfdf30d5dc890")
	productIds := make([]primitive.ObjectID, 0)
	productIds = append(productIds, productObjectId)
	var products []model.ProductRequest
	productsQuantities := make(map[primitive.ObjectID]int32, 0)
	productsQuantities[productObjectId] = 2
	products = append(products, model.ProductRequest{
		ProductId: productObjectId,
		Quantity:  2,
	})
	id, err := services.PlaceOrderProcessing(model.PlaceOrderProcessingRequest{
		ProductIds:         productIds,
		ProductsQuantities: productsQuantities,
		Products:           products,
	})

	if !utils.IsNull(id) {
		t.Fatal(`Error while placing order`, err)
	}
	fmt.Println("out of stock error ", err.Error())
}
