from dataclasses import dataclass

from catalog.domain.entities.category import Category


@dataclass
class Product:
    def __init__(self, id: int, name: str, price: float, category: Category):
        self.id = id
        self.name = name
        self.price = price
        self.category = category

    def __repr__(self) -> str:
        return f'<Product {self.id}>'
