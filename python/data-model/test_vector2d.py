from vector2d import Vector


def test_addition():
    """
    __add__ supports + operator
    """
    v1 = Vector(2, 5)
    v2 = Vector(1, 3)
    v3 = v1 + v2

    assert v1.x + v2.x == v3.x
    assert v1.y + v2.y == v3.y


def test_abs():
    """
    __abs__ supports abs method
    """
    v = Vector(3, 4)
    magnitude = abs(v)

    assert magnitude == 5


def test_mult():
    """
    __mul__ supports * operator
    """
    v = Vector(3, 4)
    result = v * 3

    assert result.x == 9
    assert result.y == 12


def test_truthy_vector():
    """
    __bool__ supports bool function
    """
    falsy_vector = Vector(0, 0)
    truthy_vector = Vector(3, 4)

    assert bool(falsy_vector) is False
    assert bool(truthy_vector) is True
