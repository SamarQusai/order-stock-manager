package context

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"simple-order-stock-manager/model/db_model"
	"time"
)

type ServiceContext interface {
	Logger() *logrus.Logger
	InitMongo() ServiceContext
	GetMongoClient() *mongo.Client
	GetDB() *mongo.Database
	WithEmailIndexes() ServiceContext
}

type Context struct {
	logger        *logrus.Logger
	mongodbClient *mongo.Client
	db            *mongo.Database
}

func NewContext() ServiceContext {
	ctx := &Context{
		logger: logrus.New(),
	}
	godotenv.Load()
	return ctx
}

func (ctx *Context) Logger() *logrus.Logger {
	return ctx.logger
}

func (ctx *Context) InitMongo() ServiceContext {
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, connectionError := mongo.Connect(contextWithTimeout, options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_URL")))
	if connectionError != nil {
		ctx.Logger().Panic("Failed to connect mongo, error:", connectionError.Error())
	}
	err := client.Ping(context.Background(), nil)
	if err != nil {
		ctx.Logger().Panic("failed to ping mongo :", err)
	}
	ctx.mongodbClient = client
	ctx.db = client.Database(os.Getenv("MONGO_DB_NAME"))
	return ctx
}

func (ctx *Context) GetDB() *mongo.Database {
	return ctx.db
}

func (ctx *Context) GetMongoClient() *mongo.Client {
	return ctx.mongodbClient
}

func (ctx *Context) WithEmailIndexes() ServiceContext {
	_, indexCreationError := ctx.GetDB().
		Collection(db_model.SentEmailCollectionName).
		Indexes().
		CreateOne(context.Background(), mongo.IndexModel{
			Keys: bson.D{{db_model.ResourceIdFieldName, 1}},
		})
	if indexCreationError != nil {
		log.Fatal("couldn't create indexes sent_email.resource_id, error: ", indexCreationError.Error())
	}
	return ctx
}
