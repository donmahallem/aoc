import sys
import typing
from collections import deque
from functools import lru_cache


def Part1(input: typing.TextIO) -> int:
    pairs = dict()
    for line in input:
        splitLine = line.strip().split(':')
        source = splitLine[0].strip()
        targets = [x.strip() for x in splitLine[1].strip().split(' ')]
        pairs[source] = targets

    @lru_cache()
    def dfs(node: str) -> int:
        if node == "out":
            return 1
        total = 0
        for successor in pairs[node]:
            total += dfs(successor)
        return total

    return dfs("you")


if __name__ == "__main__":
    print(Part1(sys.stdin))
