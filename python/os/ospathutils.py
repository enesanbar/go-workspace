import os
from os import path
import datetime
from datetime import date, time, timedelta
import time


def main():
    # Print the name of the OS
    print(f"OS Name: {os.name}")
    print(f"Current Working Directory: {os.getcwd()}")

    # Check for item existence and type
    filename = "textfile.txt"
    print("Item exists: {}".format(str(path.exists(filename))))
    print("Item is a file: " + str(path.isfile(filename)))
    print("Item is a directory: " + str(path.isdir(filename)))

    # Work with file paths
    print("Item's path: " + str(path.realpath(filename)))
    print("Item's path and name: " + str(path.split(path.realpath(filename))))

    # Get the modification time
    t = time.ctime(path.getmtime(filename))
    print(t)
    print(datetime.datetime.fromtimestamp(path.getmtime(filename)))

    # Calculate how long ago the item was modified
    td = datetime.datetime.now() - datetime.datetime.fromtimestamp(path.getmtime(filename))
    print("It has been " + str(td) + " since the file was modified")
    print("Or, " + str(td.total_seconds()) + " seconds")


if __name__ == "__main__":
    main()
