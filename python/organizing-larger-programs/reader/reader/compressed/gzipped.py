import gzip
import sys

opener = gzip.open

# utility to create gzipped text file in the command line
# python -m reader.compressed.gzipped reader-test.gz data compressed with gzip
if __name__ == '__main__':
    f = gzip.open(sys.argv[1], mode = 'wt')
    f.write(' '.join(sys.argv[2:]))
    f.close()
