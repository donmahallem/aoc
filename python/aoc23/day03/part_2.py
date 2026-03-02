import sys
import typing
from .shared import find_objects, find_adjacent_numbers


def Part2(input: typing.TextIO) -> int:
    grid = [line.rstrip("\n") for line in input if line.strip()]
    parts, numbers = find_objects(grid)

    total = 0
    for p in parts:
        if p.char != "*":
            continue
        adj = find_adjacent_numbers(p, numbers)
        if len(adj) == 2:
            total += adj[0].value * adj[1].value
    return total


if __name__ == "__main__":
    print(Part2(sys.stdin))
