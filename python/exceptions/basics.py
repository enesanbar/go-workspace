#!/usr/bin/env python3
import sys


def main():
    print('Hello, World.')
    try:
        x = 5/0
    except ValueError as e:
        print(e)
    except ZeroDivisionError as e:
        print(e)
    except Exception as e:
        print(e, sys.exc_info())
    else:
        print("no error")
        print(x)


if __name__ == '__main__':
    main()
