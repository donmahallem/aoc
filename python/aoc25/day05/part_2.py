import sys
import typing
import numpy as np
from aoc25.day05.part_1 import parseInput, compressRanges


def Part2(input: typing.TextIO) -> int:
    ranges, _ = parseInput(input)
    compresed = compressRanges(ranges)
    goodIngredients = 0
    for rL, rR in compresed:
        goodIngredients += rR - rL + 1
    return goodIngredients


if __name__ == "__main__":
    print(Part2(sys.stdin))
