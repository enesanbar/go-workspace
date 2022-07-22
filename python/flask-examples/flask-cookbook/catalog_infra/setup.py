from setuptools import find_packages, setup

setup(
    name="catalog_infra",
    version="0.0.0",
    packages=find_packages(),
    install_requires=["catalog", "injector", "redis", "sqlalchemy"],
    extras_require={},
)
