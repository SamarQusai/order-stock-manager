package model

import (
	"errors"
)

var ProductNotFound = errors.New("product not found")
var OutOfStockError = errors.New("some ingredients are out of stock")
var ProductsHaveNoIngredients = errors.New("there're products have no ingredients")
var StockNotCoversError = errors.New("stock does not cover")

var PlacingOrderErrors = []error{
	ProductNotFound,
	OutOfStockError,
	ProductsHaveNoIngredients,
	StockNotCoversError,
}
