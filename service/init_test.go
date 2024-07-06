package service

import (
	"os"
	serverContext "simple-order-stock-manager/context"
	"testing"
)

var serviceContext serverContext.ServiceContext
var services Interface

func TestMain(t *testing.M) {
	os.Setenv("MONGO_DB_NAME", "order-stock-manager-test")
	os.Setenv("MONGO_CONNECTION_URL", "mongodb://127.0.0.1:27017/?directConnection=true")
	serviceContext = serverContext.NewContext().InitMongo()
	services = New(serviceContext)
	t.Run()
}
