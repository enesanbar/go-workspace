import sys

sys.path.append('reader')
from reader import Reader

r = Reader('reader/__main__.py')
print(r.read())

r = Reader('reader/reader-test.bz2')
print(r.read())

r = Reader('reader/reader-test.gz')
print(r.read())