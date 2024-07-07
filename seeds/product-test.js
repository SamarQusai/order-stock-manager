var products = [];
    products.push(
        {
            "_id" : ObjectId("6688612be29cfdf30d5db4be"),
            "name" : "Beef Burger",
            "ingredients" : [
                {
                    "_id": ObjectId("668a6f8d133b6a4d293d1351"),
                    "name": "Beef",
                    "weight": 150,
                    "unit": "g"
                },
                {
                    "_id": ObjectId("668a6f92d4eb70c3fcbb6252"),
                    "name": "Cheese",
                    "weight": 30,
                    "unit": "g"
                },
                {
                    "_id": ObjectId("668a6f98498f41d470986ae4"),
                    "name": "Onion",
                    "weight": 20,
                    "unit": "g"
                },
            ],
        }
    )

products.push(
    {
        "_id" : ObjectId("66891e9de29cfdf30d5dc890"),
        "name" : "Chicken Burger",
        "ingredients" : [
            {
                "_id": ObjectId("668a6fd5e8fad137425f2359"),
                "name": "Chicken",
                "weight": 150,
                "unit": "g"
            },
            {
                "_id": ObjectId("668a6f92d4eb70c3fcbb6252"),
                "name": "Cheese",
                "weight": 30,
                "unit": "g"
            },
            {
                "_id": ObjectId("668a6f98498f41d470986ae4"),
                "name": "Onion",
                "weight": 20,
                "unit": "g"
            },
        ],
    }
)
use order-stock-manager-test
db.product.insertMany(products)
