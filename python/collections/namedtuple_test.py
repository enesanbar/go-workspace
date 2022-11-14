from collections import namedtuple

# Declaring namedtuple()
Student = namedtuple('Student', ['name', 'age', 'DOB'])


def test_namedtuple_access_methods():

    # Adding values
    student = Student('Nandini', '19', '2541997')

    # Access using dot notation
    assert student.name == "Nandini"
    assert getattr(student, "age") == '19'
    assert student[2] == '2541997'


def test_namedtuple_conversion_operations():

    # Adding values
    student = Student('Nandini', '19', '2541997')

    # using _make() to return namedtuple()
    print(Student._make(student))

    # using _asdict() to return an OrderedDict()
    print("The OrderedDict instance using namedtuple is  : ")
    print(student._asdict())
