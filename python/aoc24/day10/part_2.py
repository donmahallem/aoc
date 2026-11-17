import typing

from .part_1 import parseInput


def Part2(input: typing.TextIO) -> int:
    stride, size, levels = parseInput(input)
    paths = [0] * size
    other = [0] * size
    for pos in levels[9]:
        paths[pos] = 1
    for height in range(8, -1, -1):
        paths, other = other, paths
        if height < 7:
            for pos in levels[height + 3]:
                other[pos] = 0
        for pos in levels[height]:
            paths[pos] = (
                other[pos - 1]
                + other[pos + 1]
                + other[pos - stride]
                + other[pos + stride]
            )
    return sum(paths[pos] for pos in levels[0])
