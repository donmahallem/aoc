import typing
import numpy as np
from .shared import CELL_CORRUPTED, loadField

def shortestPath(gameMap: np.typing.NDArray) -> np.typing.NDArray:
    check_next = [(0, 0)]
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    pathCost = np.zeros(gameMap.shape, dtype=np.int32) - 1
    while len(check_next) > 0:
        cur_pos = check_next.pop(-1)
        cur_y, cur_x = cur_pos
        current_path_cost = max(0, pathCost[cur_y, cur_x])
        for dir_y, dir_x in dirs:
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if (
                next_y < 0
                or next_x < 0
                or next_x >= gameMap.shape[1]
                or next_y >= gameMap.shape[0]
                or gameMap[next_y, next_x] == CELL_CORRUPTED
            ):
                # OUTSIDE MAP OR CORRUPTED
                continue
            if (
                pathCost[next_y, next_x] < 0
                or current_path_cost + 1 < pathCost[next_y, next_x]
            ):
                pathCost[next_y, next_x] = current_path_cost + 1
                check_next.append((next_y, next_x))
    return pathCost

def Part1(input: typing.TextIO, size: int = 71, steps: int = 1024) -> int:
    gameMap, _ = loadField(input, size, steps)
    pathCost = shortestPath(gameMap)
    return pathCost[pathCost.shape[0] - 1, pathCost.shape[1] - 1]
