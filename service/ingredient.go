package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/model/db_model"
	"simple-order-stock-manager/utils"
)

func (c *Config) CheckIngredientsAvailability(sessionContext mongo.SessionContext, products []db_model.Product, request model.PlaceOrderProcessingRequest) ([]db_model.Ingredient, error) {
	ingredientsDetails := prepareAllIngredientsQuantity(products, request)
	if len(ingredientsDetails.Ids) == 0 {
		c.context.Logger().Error("Products have no ingredients")
		return nil, model.ProductsHaveNoIngredients
	}
	ingredientsStock, getIngredientsError := c.repository.GetIngredientsByIds(sessionContext, ingredientsDetails.Ids)
	if getIngredientsError != nil {
		c.context.Logger().Error("GetIngredientsByIds error: ", getIngredientsError.Error(), " ids:", products)
	}

	updatedIngredients, checkAvailabilityError := checkIfStockCoversAndUpdateIngredients(ingredientsDetails, ingredientsStock)
	if checkAvailabilityError != nil {
		c.context.Logger().Error("error from checkIfStockCoversAndUpdateIngredients, error: ", checkAvailabilityError.Error())
		return nil, model.OutOfStockError
	}
	return updatedIngredients, nil
}

func prepareAllIngredientsQuantity(products []db_model.Product, request model.PlaceOrderProcessingRequest) model.IngredientsDetailsPerCart {
	var ingredientsDetailsPerCart model.IngredientsDetailsPerCart
	allIngredientsInGrams := make(map[primitive.ObjectID]float64)
	var ingredientsIDs []primitive.ObjectID
	for _, product := range products {
		for _, productIngredient := range product.Ingredients {
			allIngredientsInGrams[productIngredient.ID] += productIngredient.Weight * float64(request.ProductsQuantities[product.ID])
			ingredientsIDs = append(ingredientsIDs, productIngredient.ID)
		}
	}
	ingredientsDetailsPerCart.Ids = ingredientsIDs
	ingredientsDetailsPerCart.AllIngredientsInGrams = allIngredientsInGrams
	return ingredientsDetailsPerCart
}

func checkIfStockCoversAndUpdateIngredients(cart model.IngredientsDetailsPerCart, ingredientsStock []db_model.Ingredient) ([]db_model.Ingredient, error) {

	var updatedIngredients []db_model.Ingredient
	if len(cart.AllIngredientsInGrams) != len(ingredientsStock) {
		return nil, model.OutOfStockError
	}
	for _, ingredientStock := range ingredientsStock {
		if ingredientStock.Stock < utils.ConvertFromGramToKg(cart.AllIngredientsInGrams[ingredientStock.ID]) {
			return nil, model.StockNotCoversError
		}
		ingredientStock.Stock -= utils.ConvertFromGramToKg(cart.AllIngredientsInGrams[ingredientStock.ID])
		updatedIngredients = append(updatedIngredients, ingredientStock)
	}

	return updatedIngredients, nil
}

func (c *Config) DecreaseStock(sessionContext mongo.SessionContext, ingredient db_model.Ingredient) error {
	decreaseError := c.repository.Decrease(sessionContext, ingredient.ID, ingredient.Stock)
	if decreaseError != nil {
		c.context.Logger().Error("error while decreasing stock, id: ", ingredient.ID.Hex(), " error: ", decreaseError.Error())
		return decreaseError
	}

	return nil
}

func (c *Config) notifyMerchant(ingredients []db_model.Ingredient) {
	for _, ingredient := range ingredients {
		if getIngredientAvailabilityPercentage(ingredient) < db_model.PercentageToNotifyMerchant {
			_, err := c.repository.FindEmailByResourceId(ingredient.ID)
			if err.Error() == mongo.ErrNoDocuments.Error() {
				c.context.Logger().Debug("Notify merchant with ingredient stock, ingredient id: ", ingredient.ID.Hex())
				sendingEmailError := c.sendEmail(model.SendEmailRequest{
					From:     os.Getenv("FROM_EMAIL"),
					To:       os.Getenv("MERCHANT_EMAIL"),
					Password: os.Getenv("FROM_EMAIL_PASSWORD"),
					Template: "The stock of ingredient: " + ingredient.Name + " is blew 50%",
				})
				if sendingEmailError == nil {
					c.repository.PersistEmail(db_model.SentEmail{
						EmailType:  "Stock",
						ResourceId: ingredient.ID,
					})
				}
				return
			}
		}
	}

}

func getIngredientAvailabilityPercentage(ingredient db_model.Ingredient) float64 {
	return (ingredient.Stock / ingredient.OriginalStock) * 100
}
