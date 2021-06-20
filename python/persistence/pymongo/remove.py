from pymongo import MongoClient

client = MongoClient()
db = client['test']

# removes all documents that match the specified condition.
result = db.restaurants.delete_many({"borough": "Manhattan"})

# removes all documents
result = db.restaurants.delete_many({})

# instead of removing, drop the entire collection
db.restaurants.drop()
