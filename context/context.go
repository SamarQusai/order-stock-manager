package context

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type ServiceContext interface {
	Logger() *logrus.Logger
	InitMongo() ServiceContext
	GetMongoClient() *mongo.Client
	GetDB() *mongo.Database
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
