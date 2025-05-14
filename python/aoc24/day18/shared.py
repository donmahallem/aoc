import typing
import numpy as np

NpData = np.ndarray[tuple[int, int], np.dtype[np.uint8]]


def loadField(input: typing.TextIO, size: int,
              steps: int) -> tuple[NpData, list[tuple[int, ...]]]:
    data_lines = [line.strip() for line in input.readlines()]
    data = [
        tuple(int(a) for a in line.split(","))[::-1] for line in data_lines
    ]
    gameMap = typing.cast(NpData, np.zeros((size, size), dtype=np.int8))
    for i in range(0, steps):
        y, x = data[i]
        gameMap[y, x] = CELL_CORRUPTED
    return gameMap, data


CELL_CORRUPTED = 1
