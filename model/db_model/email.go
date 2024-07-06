package db_model

import "go.mongodb.org/mongo-driver/bson/primitive"

const SentEmailCollectionName = "sent_email"

type SentEmail struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	EmailType  string             `bson:"email_type"`
	ResourceId primitive.ObjectID `bson:"resource_id"`
}
