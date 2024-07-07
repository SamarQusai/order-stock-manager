var ingredients = [];
ingredients.push(
    {
        "_id": ObjectId("668a67784111d7ca82403340"),
        "name": "Beef",
        "stock": 20,
        "unit": "kg",
        "original_stock": 20
    },
    {
        "_id": ObjectId("668a67acbe1153dde2e45a36"),
        "name": "Cheese",
        "stock": 5,
        "unit": "kg",
        "original_stock": 5

    },
    {
        "_id": ObjectId("668a67affedee0ad42e359ce"),
        "name": "Onion",
        "stock": 1,
        "unit": "kg",
        "original_stock": 1

    })

use order-stock-manager
db.ingredient.insertMany(ingredients)
