import json


class CategoryJsonEncoder(json.JSONEncoder):
    def default(self, o):
        return o.__dict__
