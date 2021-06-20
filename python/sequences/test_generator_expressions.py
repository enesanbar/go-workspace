"""
a genexp (generator expression) saves memory
because it yields items one by one using the iterator protocol
instead of building a whole list just to feed another constructor.
"""


def test_generator_expression_in_for_loop():
    colors = ['black', 'white']
    sizes = ['S', 'M', 'L']

    tshirts = tuple(f'{c} {s}' for c in colors for s in sizes)
    assert isinstance(tshirts, tuple)
    assert tshirts == ('black S', 'black M', 'black L', 'white S', 'white M', 'white L')
