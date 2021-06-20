import subprocess

cp = subprocess.run(["ls", "-lha", "asdfasdfasdf"], capture_output=True, universal_newlines=True)
print(f"stdout: {cp.stdout}")
print(f"stderr: {cp.stderr}")

# with check=True, command throws subprocess.CalledProcessError exception
try:
    cp = subprocess.run(["ls", "-lha", "asdfasdfasdf"], check=True, capture_output=True, universal_newlines=True)
except subprocess.CalledProcessError as e:
    print("An exception occurred:")
    print(e.stderr)
