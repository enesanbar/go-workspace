from defaultdict import Visits


def test_defaultdict():
    visits = Visits()

    visits.add('Spain', 'Santiago de Compostela')
    visits.add('Spain', 'Barcelona')
    visits.add('Serbia', 'Belgrade')

    assert visits.data['Spain'] == {'Barcelona', 'Santiago de Compostela'}
    assert visits.data['Serbia'] == {'Belgrade'}
    assert visits.data['Turkey'] == set()
