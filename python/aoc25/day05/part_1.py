import sys
import typing
from bisect import bisect_right


def _parseRanges(range_block: str) -> list[tuple[int, int]]:
    bounds = iter(map(int, range_block.replace("-", " ").split()))
    return list(zip(bounds, bounds))


def parseRanges(input: typing.TextIO) -> list[tuple[int, int]]:
    """Parses only the range block and skips the ingredients."""
    return _parseRanges(input.read().strip().partition("\n\n")[0])


def parseInput(input: typing.TextIO) -> tuple[list[tuple[int, int]], list[int]]:
    range_block, _, ingredient_block = input.read().strip().partition("\n\n")
    return _parseRanges(range_block), list(map(int, ingredient_block.split()))


def compressRanges(ranges: list[tuple[int, int]]) -> list[tuple[int, int]]:
    """Merges overlapping and adjacent ranges into sorted, disjoint ranges."""
    if not ranges:
        return []
    ranges = sorted(ranges)
    compressed: list[tuple[int, int]] = []
    curStart, curEnd = ranges[0]
    for start, end in ranges:
        if start <= curEnd + 1:
            curEnd = max(curEnd, end)
        else:
            compressed.append((curStart, curEnd))
            curStart, curEnd = start, end
    compressed.append((curStart, curEnd))
    return compressed


def Part1(input: typing.TextIO) -> int:
    ranges, ingredients = parseInput(input)
    compressed = compressRanges(ranges)
    starts = [start for start, _ in compressed]
    ends = [end for _, end in compressed]
    goodIngredients = 0
    for ing in ingredients:
        idx = bisect_right(starts, ing) - 1
        if idx >= 0 and ing <= ends[idx]:
            goodIngredients += 1
    return goodIngredients


if __name__ == "__main__":
    print(Part1(sys.stdin))
