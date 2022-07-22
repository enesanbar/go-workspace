import os
from dataclasses import dataclass

import injector

__all__ = ["bootstrap_app"]

from sqlalchemy import create_engine
from sqlalchemy.engine import Engine

from catalog_infra import CatalogInfrastructure
from catalog_infra.models import metadata
from main.modules import Db


@dataclass
class AppContext:
    injector: injector.Injector


def bootstrap_app() -> AppContext:
    engine = create_engine(os.environ["DB_DSN"])
    dependency_injector = _setup_dependency_injection(engine)

    _create_db_schema(engine)  # TEMPORARY

    return AppContext(dependency_injector)


def _setup_dependency_injection(engine: Engine) -> injector.Injector:
    return injector.Injector(
        [
            Db(engine),
            CatalogInfrastructure(),
        ],
        auto_bind=False,
    )


def _create_db_schema(engine: Engine) -> None:
    # Models has to be imported for metadata.create_all to discover them
    from catalog_infra import Product, Category  # noqa

    # TODO: Use migrations for that
    metadata.create_all(engine)
