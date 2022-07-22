from sqlalchemy import Column, Integer, Numeric, String, Table, MetaData, ForeignKey, Float
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship, backref

metadata = MetaData()
Base = declarative_base(metadata=metadata)


class Category(Base):
    __tablename__ = "category"

    id = Column(Integer, primary_key=True)
    name = Column(String(100))
    products = relationship('Product', back_populates='category')


class Product(Base):
    __tablename__ = "product"

    id = Column(Integer, primary_key=True)
    name = Column(String(255))
    price = Column(Float)
    category_id = Column(Integer, ForeignKey('category.id'))
    category = relationship('Category', back_populates='products')

