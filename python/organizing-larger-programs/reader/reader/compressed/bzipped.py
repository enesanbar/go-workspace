import bz2
import sys

opener = bz2.open

# utility to create bzipped text file in the command line
# python -m reader.compressed.bzipped reader-test.bz2 data compressed with bz2
if __name__ == '__main__':
    f = bz2.open(sys.argv[1], mode = 'wt')
    f.write(' '.join(sys.argv[2:]))
    f.close()
