from dataclasses import dataclass

from catalog.domain.entities.category import Category


@dataclass(frozen=True)
class CreateProductRequest:
    name: str
    price: float
    cat: Category
