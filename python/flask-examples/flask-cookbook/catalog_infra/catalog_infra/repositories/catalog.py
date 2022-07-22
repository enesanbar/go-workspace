from typing import List

from sqlalchemy.engine import Connection
from sqlalchemy.engine.row import RowProxy
from sqlalchemy.orm import Session

from catalog.application.repositories.catalog import CatalogRepository
from catalog.domain.entities import product
from catalog.domain.entities import category
from catalog.domain.requests.create_category import CreateCategoryRequest
from catalog.domain.requests.create_product import CreateProductRequest
from catalog.domain.requests.get_product import GetProductsRequest
from catalog_infra.models import Product, Category


class SqlAlchemyCatalogRepo(CatalogRepository):
    def __init__(self, connection: Connection) -> None:
        self._conn = connection

    def get_product_by_id(self, product_id: int) -> product.Product:
        with Session(self._conn.engine) as session:
            result = session.query(Product).where(Product.id == product_id).one()
            if not result:
                raise Exception("Not found")

            return self._row_to_product(result)

    def get_products(self, request: GetProductsRequest) -> List[product.Product]:
        with Session(self._conn.engine) as session:
            result = session.query(Product).offset(request.page * request.per_page).limit(request.per_page).all()

            return [self._row_to_product(row) for row in result]

    def create_product(self, r: CreateProductRequest) -> None:
        with Session(self._conn.engine) as session:
            session.add(Product(name=r.name, price=r.price, category_id=r.cat.id))
            session.commit()

    def create_category(self, request: CreateCategoryRequest) -> None:
        with Session(self._conn.engine) as session:
            session.add(Category(name=request.name))
            session.commit()
    
    def get_categories(self) -> List[category.Category]:
        with Session(self._conn.engine) as session:
            result = session.query(Category).all()

            return [self._row_to_category(row) for row in result]

    def _row_to_product(self, product_proxy: RowProxy) -> product.Product:
        return product.Product(
            product_proxy.id,
            product_proxy.name,
            product_proxy.price,
            self._row_to_category(product_proxy.category),
        )

    def _row_to_category(self, category_proxy: RowProxy) -> category.Category:
        return category.Category(
            category_proxy.id,
            category_proxy.name,
            [product.Product(id=p.id, name=p.name, price=p.price, category=category_proxy.id) for p in category_proxy.products],
        )

    def get_category_by_name(self, name: str) -> category.Category:
        with Session(self._conn.engine) as session:
            result = session.query(Category).where(Category.name == name).one()

            if not result:
                raise Exception("Not found")

            return self._row_to_category(result)
