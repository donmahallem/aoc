import sys
import typing

from aoc25.day05.part_1 import compressRanges, parseRanges


def Part2(input: typing.TextIO) -> int:
    ranges = parseRanges(input)
    return sum(end - start + 1 for start, end in compressRanges(ranges))


if __name__ == "__main__":
    print(Part2(sys.stdin))
