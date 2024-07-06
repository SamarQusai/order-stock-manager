package model

import (
	"errors"
)

var OutOfStockError = errors.New("some ingredients are out of stock")
var ProductsHaveNoIngredients = errors.New("there're products have no ingredients")
var StockNotCoversError = errors.New("stock does not cover")

var PlacingOrderErrors = []error{
	OutOfStockError,
	ProductsHaveNoIngredients,
	StockNotCoversError,
}
