var counter = 2;
var products = [];
for (i = 0; i < counter; i++) {
    products.push(
        {
            "_id" : ObjectId(),
            "name" : "Beef Burger",
            "ingredients" : [
                {
                    "_id": ObjectId("668a67784111d7ca82403340"),
                    "name": "Beef",
                    "weight": 150,
                    "unit": "g"
                },
                {
                    "_id": ObjectId("668a67acbe1153dde2e45a36"),
                    "name": "Cheese",
                    "weight": 30,
                    "unit": "g"
                },
                {
                    "_id": ObjectId("668a67affedee0ad42e359ce"),
                    "name": "Onion",
                    "weight": 20,
                    "unit": "g"
                },
            ],
        }
    )
}
use order-stock-manager
db.product.insertMany(products)
