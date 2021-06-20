from pymongo import MongoClient

client = MongoClient()
db = client['test']

# Update the first document with name equal to "Juni"
# result contains result.matched_count and result.modified_count
result = db.restaurants.update_one(
    {"name": "Juni"},
    {
        "$set": {
            "cuisine": "American (New)"
        },
        "$currentDate": {
            "lastModified": True
        }
    }
)

# Update multiple documents
result = db.restaurants.update_many(
    {"address.zipcode": "10016", "cuisine": "Other"},
    {
        "$set": {"cuisine": "Category to be determined"},
        "$currentDate": {"lastModified": True}
    }
)

# Replace an entire document
result = db.restaurants.replace_one(
    {"restaurant_id": "41704620"},
    {
        "name": "Vella 2",
        "address": {
            "coord": [-73.9557413, 40.7720266],
            "building": "1480",
            "street": "2 Avenue",
            "zipcode": "10075"
        }
    }
)
