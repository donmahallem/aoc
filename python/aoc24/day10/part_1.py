import typing


def parseInput(input: typing.TextIO) -> tuple[int, int, list[list[int]]]:
    lines = input.read().split()
    stride = len(lines[0]) + 1
    grid = ".".join(lines) + "." * (stride + 1)
    levels: list[list[int]] = [[] for _ in range(11)]
    append = {char: level.append for char, level in zip("0123456789.", levels)}
    for pos, char in enumerate(grid):
        append[char](pos)
    return stride, len(grid), levels


def Part1(input: typing.TextIO) -> int:
    stride, size, levels = parseInput(input)
    reach = [0] * size
    other = [0] * size
    for bit, pos in enumerate(levels[9]):
        reach[pos] = 1 << bit
    for height in range(8, -1, -1):
        reach, other = other, reach
        if height < 7:
            for pos in levels[height + 3]:
                other[pos] = 0
        for pos in levels[height]:
            reach[pos] = (
                other[pos - 1]
                | other[pos + 1]
                | other[pos - stride]
                | other[pos + stride]
            )
    return sum(reach[pos].bit_count() for pos in levels[0])
