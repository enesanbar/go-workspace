import iso6346


class ShippingContainer:
    # class attributes
    next_serial = 1335
    HEIGHT_FT = 8.5
    WIDTH_FT = 8.0

    @staticmethod
    def _make_bic_code(owner_code, serial):
        return iso6346.create(owner_code = owner_code,
                              serial = str(serial).zfill(6))

    @classmethod
    def create_empty(cls, owner_code, length_ft, *args, **kwargs):
        return cls(owner_code, length_ft, contents = None, *args, **kwargs)

    @classmethod
    def create_with_items(cls, owner_code, length_ft, items, *args, **kwargs):
        return cls(owner_code, length_ft, list(items), *args, **kwargs)

    def __init__(self, owner_code, length_ft, contents):
        # instance attributes
        self.owner_code = owner_code
        self.contents = contents
        self.length_ft = length_ft
        self.bic = self._make_bic_code(owner_code = owner_code,
                                       serial = ShippingContainer._get_next_serial())

    @property
    def volume_ft3(self):
        return self._calculate_volume()

    def _calculate_volume(self):
        return ShippingContainer.HEIGHT_FT * ShippingContainer.WIDTH_FT * self.length_ft

    @classmethod
    def _get_next_serial(cls):
        result = cls.next_serial
        cls.next_serial += 1
        return result

    # used for debugging and logging purposes and contains much more information than __str__
    def __repr__(self):
        return "ShippingContainer(owner_code='{}', length_ft={}, contents=list({}))".format(self.owner_code,
                                                                                            self.length_ft,
                                                                                            self.contents)

    # used when object is printed with print() for a readable, human-friendly representation
    def __str__(self):
        return "BIC: {}".format(self.bic)

    # used when object is printed with '{}'.format(obj)
    # '{:option}'.format(): option is passed to __format__
    # '{!r}'.format(): use __repr__ representation instead of __format__
    # '{!s}'.format(): use __str__ representation instead of __format__
    def __format__(self, format_spec):
        return self.bic


class RefrigeratedShippingContainer(ShippingContainer):
    MAX_CELSIUS = 4.0

    FRIDGE_VOLUME_FT3 = 100

    @staticmethod
    def _make_bic_code(owner_code, serial):
        return iso6346.create(owner_code = owner_code,
                              serial = str(serial).zfill(6),
                              category = 'R')

    @staticmethod
    def _celsius_to_fahrenheit(celsius):
        return celsius * 9 / 5 + 32

    @staticmethod
    def _fahrenheit_to_celsius(fahrenheit):
        return (fahrenheit - 32) * 5 / 9

    def __init__(self, owner_code, length_ft, contents, celsius):
        super().__init__(owner_code, length_ft, contents)
        self.celsius = celsius

    @property
    def celsius(self):
        return self._celsius

    @celsius.setter
    def celsius(self, value):
        self._set_celsius(value)

    def _set_celsius(self, value):
        if value > RefrigeratedShippingContainer.MAX_CELSIUS:
            raise ValueError('Temperature too hot!')
        self._celsius = value

    @property
    def fahrenheit(self):
        return RefrigeratedShippingContainer._celsius_to_fahrenheit(self.celsius)

    @fahrenheit.setter
    def fahrenheit(self, value):
        self.celsius = RefrigeratedShippingContainer._fahrenheit_to_celsius(value)

    def _calculate_volume(self):
        return super()._calculate_volume() - RefrigeratedShippingContainer.FRIDGE_VOLUME_FT3


class HeatedRefrigeratedShippingContainer(RefrigeratedShippingContainer):
    MIN_CELSIUS = -20.0

    def _set_celsius(self, value):
        if value < HeatedRefrigeratedShippingContainer.MIN_CELSIUS:
            raise ValueError('Temperature too cold!')
        super()._set_celsius(value)
