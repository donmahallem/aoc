import sys
import typing
from functools import lru_cache
from .parse_input import __parseInput

Coord: typing.TypeAlias = tuple[int, int]


def Part2(input: typing.TextIO) -> int:
    result = __parseInput(input)
    (x, y), splitter, width, height = result

    cache: dict[Coord, int] = {}

    def dfs(bx: int, by: int) -> int:
        key: Coord = (bx, by)

        if key in cache:
            return cache[key]

        if by == height:
            return 1

        total = 0
        beam_pos = bx + (by * width)

        if beam_pos in splitter:
            if bx > 0:
                total += dfs(bx - 1, by + 1)
            if bx < width - 1:
                total += dfs(bx + 1, by + 1)
        else:
            total += dfs(bx, by + 1)

        cache[key] = total
        return total

    return dfs(x, y)


if __name__ == "__main__":
    print(Part2(sys.stdin))
