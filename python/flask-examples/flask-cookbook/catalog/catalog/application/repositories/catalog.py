import abc
from typing import List

from catalog.domain.entities.product import Product, Category
from catalog.domain.requests.create_category import CreateCategoryRequest
from catalog.domain.requests.create_product import CreateProductRequest
from catalog.domain.requests.get_product import GetProductsRequest


class CatalogRepository(abc.ABC):

    @abc.abstractmethod
    def get_product_by_id(self, product_id: int) -> Product:
        pass

    @abc.abstractmethod
    def get_products(self, request: GetProductsRequest) -> List[Product]:
        pass

    @abc.abstractmethod
    def create_product(self, request: CreateProductRequest) -> None:
        pass

    @abc.abstractmethod
    def get_category_by_name(self, name: str) -> Category:
        pass

    @abc.abstractmethod
    def create_category(self, request: CreateCategoryRequest) -> None:
        pass

    @abc.abstractmethod
    def get_categories(self) -> List[Category]:
        pass
