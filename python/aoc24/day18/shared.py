import typing
import numpy as np


def loadField(
    input: typing.TextIO, size: int, steps: int
) -> tuple[np.typing.NDArray, list[tuple[int, int]]]:
    data = [line.strip() for line in input.readlines()]
    data = [tuple(int(a) for a in line.split(","))[::-1] for line in data]
    gameMap = np.zeros((size, size), dtype=np.int8)
    for i in range(0, steps):
        y, x = data[i]
        gameMap[y, x] = CELL_CORRUPTED
    return gameMap, data


CELL_CORRUPTED = 1
