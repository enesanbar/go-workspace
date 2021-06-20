import os
from reader.compressed import *

extension_map = {
    '.bz2': bz2_opener,
    '.gz': gzip_opener
}


class Reader:
    def __init__(self, filename):
        _, extension = os.path.splitext(filename)
        opener = extension_map.get(extension, open)
        self.file = opener(filename, 'rt')

    def close(self):
        self.file.close()

    def read(self):
        return self.file.read()
