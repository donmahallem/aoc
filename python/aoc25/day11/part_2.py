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
    def dfs(currentNode: str, targetNode: str) -> int:
        if currentNode == targetNode:
            return 1
        total = 0
        if currentNode not in pairs:
            return 0
        for successor in pairs[currentNode]:
            total += dfs(successor, targetNode)
        return total

    path1 = dfs("svr", "dac") * dfs("dac", "fft") * dfs("fft", "out")
    path2 = dfs("svr", "fft") * dfs("fft", "dac") * dfs("dac", "out")

    return path1 + path2


if __name__ == "__main__":
    print(Part2(sys.stdin))
