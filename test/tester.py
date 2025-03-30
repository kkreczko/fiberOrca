# RUN FROM TERMINAL TO HIDE PASSWORD
import multiprocessing
import subprocess
import random

BACKEND_PATH = "../release/network_analyzer"

HOSTS = [
    "8.8.8.8",
    "8.8.4.4",
    "www.wp.pl",
    "www.google.com",
    "www.onet.pl",
    "www.interia.pl",
    "www.allegro.pl",
    "www.youtube.com",
    "www.facebook.com",
    "www.amazon.com",
    "www.linkedin.com",
    "www.github.com",
    "www.microsoft.com",
    "www.apple.com",
    "www.chess.com",
    "www.wikipedia.org",
    "www.yahoo.com",
    "www.x.com"
]

# RUN THIS CODE FROM TERMINAL TO HIDE PASSWORD
class TestCase:
    def __init__(self, packets: int, host: str) -> None:
        self.packets = packets
        self.host = host
        self.passed = False

    def run_test(self) -> None:
        run_command = ["sudo", BACKEND_PATH, "-n", str(self.packets), "host", self.host, "-t"]
        ping_command = ["ping", "-c", str(self.packets), self.host]
        try:
            #it's stupid but it works in our case
            subprocess.run(["sudo", "ls"], capture_output=True)
            result = subprocess.Popen(run_command,
                                      stdout=subprocess.DEVNULL,
                                      stderr=subprocess.DEVNULL)
            pinger = subprocess.Popen(ping_command,
                                      stdout=subprocess.DEVNULL,
                                      stderr=subprocess.DEVNULL)
            result.wait()
            pinger.wait()
            self.passed = result.returncode == 0
        except Exception as e:
            print(e)

    def __str__(self) -> str:
        return f"Packets: {self.packets}, Host: {self.host}"

def worker(test_case: TestCase) -> TestCase:
    test_case.run_test()
    return test_case

# RUN FROM TERMINAL TO HIDE PASSWORD
if __name__ == '__main__':
    test_cases = [TestCase(random.randint(2, 8), item) for item in HOSTS]

    results = []
    try:
        with multiprocessing.Pool() as pool:
            results = pool.map(worker, test_cases)
    except Exception as e:
        print(e)

    passed = len([item for item in results if item.passed])
    print(f"Total: {len(test_cases)}, Passed: {passed}, Failed: {len(test_cases) - passed}, Success rate: {passed / len(test_cases) * 100:.2f}%")
    for item in results:
        print(f"{item.packets} packets at {item.host} {"PASSED" if item.passed else "FAILED"}")

