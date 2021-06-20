# Checking class types and instances
class Book:
    def __init__(self, title):
        self.title = title


class Newspaper:
    def __init__(self, name):
        self.name = name


# Create some instances of the classes
b1 = Book("The Catcher In The Rye")
b2 = Book("The Grapes of Wrath")
n1 = Newspaper("The Washington Post")
n2 = Newspaper("The New York Times")

print(f'type(b1): {type(b1)}')
print(f'type(n1): {type(n1)}')

print(f'type(b1) == type(b2): {type(b1) == type(b2)}')
print(f'type(b1) == type(n2): {type(b1) == type(n2)}')

print(isinstance(b1, Book))
print(isinstance(n1, Newspaper))

print(isinstance(n2, Book))
print(isinstance(n2, object))
