import sys
import typing
from .shared import find_objects, find_adjacent_numbers


def Part1(input: typing.TextIO) -> int:
    grid = [line.rstrip("\n") for line in input if line.strip()]
    parts, numbers = find_objects(grid)

    # collect all numbers that are adjacent to at least one part
    adjacent_set: set[int] = set()
    for p in parts:
        for num in find_adjacent_numbers(p, numbers):
            adjacent_set.add(id(num))

    return sum(n.value for n in numbers if id(n) in adjacent_set)


if __name__ == "__main__":
    print(Part1(sys.stdin))
