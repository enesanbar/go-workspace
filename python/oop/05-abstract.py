# Using Abstract Base Classes to enforce class constraints
import json
from abc import ABC, abstractmethod


class GraphicShape(ABC):
    # Inheriting from ABC indicates that this is an abstract base class
    def __init__(self):
        super().__init__()

    # declaring a method as abstract requires a subclass to implement it
    @abstractmethod
    def calc_area(self):
        pass


class JSONify(ABC):
    @abstractmethod
    def to_json(self):
        pass


class Circle(GraphicShape, JSONify):
    def __init__(self, radius):
        super().__init__()
        self.radius = radius

    def calc_area(self):
        return 3.14 * (self.radius ** 2)

    def to_json(self) -> str:
        return json.dumps({'radius': self.radius, 'area': self.calc_area()})


class Square(GraphicShape, JSONify):
    def __init__(self, side):
        super().__init__()
        self.side = side

    def calc_area(self):
        return self.side * self.side

    def to_json(self) -> str:
        return json.dumps({'side': self.side, 'area': self.calc_area()})


# Abstract classes can't be instantiated themselves
# g = GraphicShape() # this will error

c = Circle(10)
print(c.calc_area())
print(c.to_json())
s = Square(12)
print(s.calc_area())
print(s.to_json())
