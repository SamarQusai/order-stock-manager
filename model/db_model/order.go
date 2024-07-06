package db_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const OrderCollectionName = "order"

type Order struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Items []Item             `json:"items" bson:"items"`
}

type Item struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Price float64            `json:"price" bson:"price"`
}
