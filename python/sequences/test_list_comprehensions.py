def test_list_comprehensions():
    languages = ["python", "java", "go", "javascript"]
    uppercase_languages = [lang.upper() for lang in languages]

    assert uppercase_languages == ["PYTHON", "JAVA", "GO", "JAVASCRIPT"]


def test_list_comprehension_scope_with_walrus_operator():
    languages = ["python", "java", "go", "javascript"]
    uppercase_languages = [last := lang.upper() for lang in languages]

    assert last == uppercase_languages[-1]
