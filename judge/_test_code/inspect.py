import os
import sys
import subprocess

def run_cmd(cmd: list[str]) -> None:
    print(f"==== executing command ({cmd=}) ====", flush=True)
    subprocess.run(cmd)
    print("===========================")

def main():
    print(f"{sys.version_info=}")
    print(f"{os.getenv('USER')=}")
    print(f"{os.getenv('UID')=}")
    print(f"{os.getenv('GID')=}")
    print(f"{os.getenv('SZPP_JUDGE')=}")
    sys.stdout.flush()

    run_cmd(["id"])
    run_cmd(["pwd"])
    run_cmd(["ls",  "-la", "/", "/var"])
    run_cmd(["bash", "-c", "ulimit -a"])

main()
