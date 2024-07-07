package repoistory

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-order-stock-manager/model/db_model"
)

func (c *Config) FindEmailByResourceId(resourceId primitive.ObjectID) (*db_model.SentEmail, error) {
	var sentEmailObject db_model.SentEmail
	findEmailError := c.db.Collection(db_model.SentEmailCollectionName).FindOne(context.Background(), bson.M{"resource_id": resourceId}).Decode(&sentEmailObject)
	if findEmailError != nil {
		if findEmailError.Error() == mongo.ErrNoDocuments.Error() {
			c.context.Logger().Debug("No email sent for resource id: ", resourceId)
			return nil, findEmailError
		}

		c.context.Logger().Error("Error while finding email for resource id: ", resourceId, " error: ", findEmailError.Error())
		return nil, findEmailError
	}

	return &sentEmailObject, nil
}

func (c *Config) PersistEmail(email db_model.SentEmail) error {
	_, insertingError := c.db.Collection(db_model.SentEmailCollectionName).InsertOne(context.Background(), email)
	if insertingError != nil {
		c.context.Logger().Error("Error while persisting email, error: ", insertingError.Error())
		return insertingError
	}

	return nil

}
