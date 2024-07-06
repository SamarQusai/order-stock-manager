package repoistory

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) GetIngredientById(sessionContext mongo.SessionContext, id primitive.ObjectID) (*db_model.Ingredient, error) {
	var ingredientObject db_model.Ingredient
	findIngredientError := c.db.Collection(db_model.IngredientCollectionName).FindOne(sessionContext, bson.M{"_id": id}).Decode(&ingredientObject)
	if findIngredientError != nil {
		c.context.Logger().Error("Error while finding product ingredient, id: ", id, " error: ", findIngredientError.Error())
		if findIngredientError.Error() == mongo.ErrNoDocuments.Error() {
			return nil, model.OutOfStockError
		}
		return nil, findIngredientError
	}

	return &ingredientObject, nil
}

func (c *Config) GetIngredientsByIds(sessionContext mongo.SessionContext, ids []primitive.ObjectID) ([]db_model.Ingredient, error) {
	var ingredientsObject []db_model.Ingredient
	cursor, findIngredientsError := c.db.Collection(db_model.IngredientCollectionName).Find(sessionContext, bson.M{"_id": bson.M{"$in": ids}})
	if findIngredientsError != nil {
		c.context.Logger().Error("Error while finding ingredients, error: ", findIngredientsError.Error())
	}
	decodingError := cursor.All(context.Background(), &ingredientsObject)
	if decodingError != nil {
		c.context.Logger().Error("Decoding error while finding ingredients - decodingErr, ", decodingError.Error())
		return nil, decodingError
	}
	return ingredientsObject, nil
}

func (c *Config) Decrease(sessionContext mongo.SessionContext, id primitive.ObjectID, stock float64) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"stock": stock}}
	_, err := c.db.Collection(db_model.IngredientCollectionName).UpdateOne(sessionContext, filter, update)
	if err != nil {
		c.context.Logger().Error(" Failed to decrease stock for ingredient: ", id, " error: ", err.Error())
		return err
	}
	return nil
}
