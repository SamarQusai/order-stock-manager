package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"simple-order-stock-manager/model/db_model"
)

type PlaceOrderProcessingRequest struct {
	ProductIds         []primitive.ObjectID
	ProductsQuantities map[primitive.ObjectID]int32
	Products           []ProductRequest
}

type PlaceOrderProcessingResponse struct {
	OrderId     string
	Ingredients []db_model.Ingredient
}

type ProductRequest struct {
	ProductId primitive.ObjectID
	Quantity  int32
}

type SendEmailRequest struct {
	From     string
	To       string
	Password string
	Template string
}

type IngredientsDetailsPerCart struct {
	Ids                   []primitive.ObjectID
	AllIngredientsInGrams map[primitive.ObjectID]float64
}
