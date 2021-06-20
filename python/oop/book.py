import warnings

# With extracted author class
class Author:
    def __init__(self, author_data):  # <1>
        self.first_name = author_data['first_name']
        self.last_name = author_data['last_name']

    @property
    def for_display(self):  # <2>
        return f'{self.first_name} {self.last_name}'

    @property
    def for_citation(self):
        return f'{self.last_name}, {self.first_name[0]}.'


class Book:
    def __init__(self, data):
        # ...

        self.author_data = data['author']  # <3>
        self.author = Author(self.author_data)  # <4>

    @property
    def author_for_display(self):  # <5>
        warnings.warn('Use book.author.for_display instead', DeprecationWarning)
        return self.author.for_display

    @property
    def author_for_citation(self):
        warnings.warn('Use book.author.for_citation instead', DeprecationWarning)
        return self.author.for_citation
