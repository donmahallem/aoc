import sys
import typing
from collections import deque
from functools import lru_cache


def Part2(input: typing.TextIO) -> int:
    pairs = dict()
    for line in input:
        line = line.strip().split(':')
        source = line[0].strip()
        targets = [x.strip() for x in line[1].strip().split(' ')]
        pairs[source] = targets

    @lru_cache()
    def dfs(currentNode: str,
            targetNode: str,
            exitOnDac: bool = False,
            exitOnFft: bool = False) -> int:
        if currentNode == targetNode:
            return 1
        total = 0
        if currentNode not in pairs:
            return 0
        for successor in pairs[currentNode]:
            if exitOnDac and successor == "dac":
                continue
            if exitOnFft and successor == "fft":
                continue
            total += dfs(successor, targetNode, exitOnDac, exitOnFft)
        return total

    path1 = dfs("svr", "dac", False, True) * dfs(
        "dac", "fft", False, False) * dfs("fft", "out", False, False)
    path2 = dfs("svr", "fft", True, False) * dfs(
        "fft", "dac", False, False) * dfs("dac", "out", False, False)

    return path1 + path2


if __name__ == "__main__":
    print(Part2(sys.stdin))
