import typing

CELL_CORRUPTED = 1


def loadField(
    input: typing.TextIO, size: int, steps: int
) -> tuple[list[list[int]], list[tuple[int, ...]]]:
    data_lines = [line.strip() for line in input.readlines()]
    data = [tuple(int(a) for a in line.split(","))[::-1] for line in data_lines]
    gameMap: list[list[int]] = [[0] * size for _ in range(size)]
    for i in range(steps):
        y, x = data[i]
        gameMap[y][x] = CELL_CORRUPTED
    return gameMap, data


CELL_CORRUPTED = 1
