# Using composition to build complex objects

class Book:
    def __init__(self, title, price, author=None):
        self.title = title
        self.price = price

        # Use references to other objects, like author and chapters
        self.author = author
        self.chapters = []

    def add_chapter(self, chapter):
        self.chapters.append(chapter)

    def get_page_count(self):
        result = 0
        for ch in self.chapters:
            result += ch.page_count
        return result


class Author:
    def __init__(self, fname, lname):
        self.fname = fname
        self.lname = lname

    def __str__(self):
        return f"{self.fname} {self.lname}"


class Chapter:
    def __init__(self, name, page_count):
        self.name = name
        self.page_count = page_count


auth = Author("Leo", "Tolstoy")
b1 = Book("War and Peace", 39.95, auth)

b1.add_chapter(Chapter("Chapter 1", 104))
b1.add_chapter(Chapter("Chapter 2", 89))
b1.add_chapter(Chapter("Chapter 3", 124))

print(b1.title)
print(b1.author)
print(b1.get_page_count())
