import pytest

from tool import Tool

tools = [
    Tool('level', 3.5),
    Tool('hammer', 1.25),
    Tool('screwdriver', 0.5),
    Tool('chisel', 0.55),
]


def test_sorting_does_not_work_without_special_dunder_methods():
    with pytest.raises(TypeError):
        tools.sort()


def test_sorting_with_lambda_by_object_attribute():
    print('Unsorted:', repr(tools))
    tools.sort(key=lambda x: x.name)
    print('\nSorted by name: ', tools)
    assert tools[0].name == 'chisel'

    tools.sort(key=lambda x: x.weight)
    print('\nSorted by weight: ', tools)
    assert tools[0].name == 'screwdriver'


def test_sorting_with_lambda_by_object_attribute_with_transformation():
    places = ['home', 'work', 'New York', 'Paris']
    places.sort()
    print('Case sensitive: ', places)
    assert places[0] == 'New York'
    assert places[-1] == 'work'

    places.sort(key=lambda x: x.lower())
    print('Case insensitive: ', places)
    assert places[0] == 'home'
    assert places[-1] == 'work'


def test_sorting_by_multiple_criteria_with_tuples():
    power_tools = [
        Tool('drill', 4),
        Tool('circular saw', 5),
        Tool('jackhammer', 40),
        Tool('sander', 4),
    ]

    print(f'Unsorted: {power_tools}')
    power_tools.sort(key=lambda x: (x.weight, x.name))
    print(f'Sorted by weight and name: {power_tools}')
    assert power_tools[0].name == 'drill'
    assert power_tools[-1].name == 'jackhammer'
