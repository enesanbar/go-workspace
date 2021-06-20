from frenchdeck import Card, FrenchDeck


def test_card():
    beer_card = Card('7', 'diamonds')  # Card(rank='7', suit='diamonds')
    assert beer_card.rank == '7'
    assert beer_card.suit == 'diamonds'


def test_deck_length():
    """
    __len__ method

    a deck responds to the len() function by returning the number of cards in it
    """
    deck = FrenchDeck()
    assert len(deck) == 52


def test_get_item():
    """
    __getitem__ method enables indexing, slicing, iteration, search
    """
    deck = FrenchDeck()

    # indexing
    first_item = deck[0]  # Card(rank='2', suit='spades')
    last_item = deck[-1]  # Card(rank='A', suit='hearts')

    assert first_item.rank == '2' and first_item.suit == 'spades'
    assert last_item.rank == 'A' and last_item.suit == 'hearts'

    # slicing
    # [Card(rank='2', suit='spades'), Card(rank='3', suit='spades'), Card(rank='4', suit='spades')]
    cards = deck[:3]
    assert len(cards) == 3

    # iteration
    for card in deck:
        print(card)

    # searching with 'in'
    assert Card('Q', 'hearts') in deck
    assert Card('7', 'beast') not in deck

    # sorting
    suit_values = dict(spades=3, hearts=2, diamonds=1, clubs=0)

    def spades_high(card):
        rank_value = FrenchDeck.ranks.index(card.rank)
        return rank_value * len(suit_values) + suit_values[card.suit]

    sorted_deck = sorted(deck, key=spades_high)
    assert sorted_deck[0].rank == '2' and sorted_deck[0].suit == 'clubs'
    assert sorted_deck[-1].rank == 'A' and sorted_deck[-1].suit == 'spades'
