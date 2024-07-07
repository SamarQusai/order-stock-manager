var ingredients = [];
ingredients.push(
    {
        "_id": ObjectId("668a6f8d133b6a4d293d1351"),
        "name": "Beef",
        "stock": 20,
        "unit": "kg",
        "original_stock": 20
    },
    {
        "_id": ObjectId("668a6f92d4eb70c3fcbb6252"),
        "name": "Cheese",
        "stock": 5,
        "unit": "kg",
        "original_stock": 5

    },
    {
        "_id": ObjectId("668a6f98498f41d470986ae4"),
        "name": "Onion",
        "stock": 1,
        "unit": "kg",
        "original_stock": 1

    },
    {
        "_id": ObjectId("668a6fd5e8fad137425f2359"),
        "name": "Chicken",
        "stock": .09,
        "unit": "kg",
        "original_stock": 1

    })

use order-stock-manager-test
db.ingredient.insertMany(ingredients)
