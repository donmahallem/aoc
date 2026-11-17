import sys
import typing

from aoc25.day04.part_1 import accessible, parseInput


def Part2(input: typing.TextIO) -> int:
    grid, stride = parseInput(input)
    removed = 0
    # A removal only makes other rolls more accessible, so removing all
    # accessible rolls per round gives the same total as one at a time.
    while movable := accessible(grid, stride):
        removed += movable.bit_count()
        grid ^= movable
    return removed


if __name__ == "__main__":
    print(Part2(sys.stdin))
