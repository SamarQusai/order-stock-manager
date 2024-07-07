package main

import (
	defaultContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"simple-order-stock-manager/context"
	"simple-order-stock-manager/server"
	"time"
)

func main() {
	serviceContext := context.NewContext().InitMongo().WithEmailIndexes()
	defer func() {
		contextWithTimeout, cancel := defaultContext.WithTimeout(defaultContext.Background(), 10*time.Second)
		defer cancel()
		if disconnectError := serviceContext.GetMongoClient().Disconnect(contextWithTimeout); disconnectError != nil {
			serviceContext.Logger().Panic("Failed to disconnect mongo, error:", disconnectError.Error())
		}
	}()
	serviceContext.Logger().Info("Starting order stock manager server")

	router := server.New(serviceContext)
	engine := gin.Default()
	router.Install(engine)
	runningError := engine.Run(":" + os.Getenv("PORT"))
	if runningError != nil {
		log.Fatal("Failed to run server", runningError.Error())
	}
}
