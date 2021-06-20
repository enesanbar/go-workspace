import os


def find_rc(rc_name='.examplerc'):
    # Check for env variable
    var_name = 'EXAMPLERC_DIR'
    if var_name in os.environ:
        var_path = os.path.join(f"${var_name}", rc_name)
        config_path = os.path.expandvars(var_path)
        print(f"Checking {config_path}")
        if os.path.exists(config_path):
            return config_path

    # Check current working directory
    config_path = os.path.join(os.getcwd(), rc_name)
    print(f'Checking {config_path}')
    if os.path.exists(config_path):
        return config_path

    # Check user home directory
    home_dir = os.path.expanduser('~/')
    config_path = os.path.join(home_dir, rc_name)
    print(f'Checking {config_path}')
    if os.path.exists(config_path):
        return config_path

    # Check directory of this file
    file_path = os.path.abspath(__file__)
    parent_path = os.path.dirname(file_path)
    config_path = os.path.join(parent_path, rc_name)
    print(f'Checking {config_path}')
    if os.path.exists(config_path):
        return config_path

    print(f"File {rc_name} has not been found")


def main():
    print(find_rc('.zshrc'))


if __name__ == '__main__':
    main()

