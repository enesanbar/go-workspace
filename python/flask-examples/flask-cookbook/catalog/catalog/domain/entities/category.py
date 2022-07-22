from dataclasses import dataclass


@dataclass
class Category:
    def __init__(self, id: int, name: str, products):
        self.id = id
        self.name = name
        self.products = products

    def __repr__(self) -> str:
        return f'<Category {self.id}>'
