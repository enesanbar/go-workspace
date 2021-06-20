#!/usr/bin/env python
"""
Simple command-line tool using sys.argv
"""

import sys

if __name__ == '__main__':
    for idx, arg in enumerate(sys.argv):
        print(f'[{idx}] {arg}')
