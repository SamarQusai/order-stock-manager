package db_model

import "go.mongodb.org/mongo-driver/bson/primitive"

//docker exec -it mongop mongosh --eval "rs.initiate({ _id: "myReplicaSet", members: [ {_id: 0, host: "mongop"}] })"

const ProductCollectionName = "product"

type Product struct {
	ID          primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Ingredients []ProductIngredient `json:"ingredients" bson:"ingredients"`
}

type ProductIngredient struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Weight float64            `json:"weight" bson:"weight"`
	Unit   string             `json:"unit" bson:"unit"`
}
