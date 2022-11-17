def test_prefer_join_over_concat():
    first_name = 'Enes'
    last_name = 'Anbar'

    # Not a recommended way to concatenate string
    full_name = first_name + "  " + last_name

    # More performant and improve readability
    full_name = " ".join([first_name, last_name])

    assert full_name == 'Enes Anbar'
