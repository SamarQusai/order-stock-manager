package db_model

import "go.mongodb.org/mongo-driver/bson/primitive"

const IngredientCollectionName = "ingredient"

const (
	KGUnit                     = "kg"
	PercentageToNotifyMerchant = 50
)

type Ingredient struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Stock         float64            `json:"stock" bson:"stock"`
	Unit          string             `json:"unit" bson:"unit"`
	OriginalStock float64            `json:"original_stock" bson:"original_stock"`
}
