import injector
from sqlalchemy.engine import Connection

from catalog.application.cache.catalog import CatalogCache
from catalog.application.repositories.catalog import CatalogRepository
from catalog_infra.cache.catalog import RedisCache
from catalog_infra.models import Product, Category
from catalog_infra.repositories import SqlAlchemyCatalogRepo

__all__ = [
    # module
    "CatalogInfrastructure",
    # models
    "Product",
    "Category",
]


class CatalogInfrastructure(injector.Module):

    @injector.provider
    def catalog_repo(self, conn: Connection) -> CatalogRepository:
        return SqlAlchemyCatalogRepo(conn)

    @injector.provider
    def catalog_cache(self) -> CatalogCache:
        return RedisCache()
