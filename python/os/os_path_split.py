import os


def main():
    current_dir = os.getcwd()
    print(f"Current dir: {current_dir}")
    print(os.path.split(current_dir))
    print(f"Parent dir name: {os.path.dirname(current_dir)}")
    print(f"Current dir name: {os.path.basename(current_dir)}")


if __name__ == '__main__':
    main()
