import subprocess

n = int(input())
for i in range(n):
    subprocess.Popen(["sleep", "10"])
    print(f"spawned {i+1}/{n}", flush=True)


subprocess.run(["bash", "-c", "ls -l /proc/*/exe"])
