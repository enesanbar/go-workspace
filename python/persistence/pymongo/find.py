import pymongo
from pymongo import MongoClient

client = MongoClient()
db = client.test

# Query for All Documents in a Collection
cursor = db.restaurants.find()

# Query for all documents in a collection whose 'borough' field is equal to 'Manhattan'
cursor = db.restaurants.find({"borough": "Manhattan"})

# Query by a field in an embedded document
cursor = db.restaurants.find({"address.zipcode": "10075"})

# Query by a field in an array
cursor = db.restaurants.find({"grades.grade": "B"})

# Greater than and less than operator
cursor = db.restaurants.find({"grades.score": {"$gt": 30}})
cursor = db.restaurants.find({"grades.score": {"$lt": 10}})

# Logical AND operator
cursor = db.restaurants.find({"cuisine": "Italian", "address.zipcode": "10075"})

# Logical OR operator
cursor = db.restaurants.find({
    "$or": [{"cuisine": "Italian"}, {"address.zipcode": "10075"}]
})

# Sorting the query result
cursor = db.restaurants.find().sort([
    ("borough", pymongo.ASCENDING),
    ("address", pymongo.ASCENDING)
])

for document in cursor:
    print document