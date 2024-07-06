package server

import (
	"github.com/gin-gonic/gin"
	"simple-order-stock-manager/context"
	"simple-order-stock-manager/service"
)

type Config struct {
	context  context.ServiceContext
	services service.Interface
}

func New(context context.ServiceContext) *Config {
	return &Config{
		context:  context,
		services: service.New(context),
	}
}

func (c *Config) Install(engine *gin.Engine) {
	orderApis := engine.Group("orders")
	{
		orderApis.POST("", c.placeOrder)
	}
}
