def test_prefer_isinstance_over_type():
    """
    because isinstance() is true for subclasses
    """
    user_ages = {"Larry": 35, "Jon": 89, "Imli": 12}
    assert type(user_ages) == dict
    assert isinstance(user_ages, dict)

    class UpperCaseDict(dict):
        def __setitem__(self, key, value):
            key = key.upper()
            super().__setitem__(key, value)

    numbers = UpperCaseDict()
    assert not type(numbers) == dict
    assert isinstance(numbers, dict)
