import typing
import numpy as np
from .shared import CELL_CORRUPTED, loadField

def shortestPath(gameMap: np.typing.NDArray):
    gameMapWidth = gameMap.shape[1]
    gameMapHeight = gameMap.shape[0]
    check_next = [(0, 0)]
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    path_cost = np.zeros(gameMap.shape, dtype=np.int32) - 1
    path_cost[0, 0] = 0
    while len(check_next) > 0:
        cur_pos = check_next.pop(-1)
        cur_y, cur_x = cur_pos
        current_path_cost = max(0, path_cost[cur_y, cur_x])
        for dir_y, dir_x in dirs:
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if (
                next_y < 0
                or next_x < 0
                or next_x >= gameMapWidth
                or next_y >= gameMapHeight
                or gameMap[next_y, next_x] == CELL_CORRUPTED
            ):
                continue
            if (
                path_cost[next_y, next_x] < 0
                or current_path_cost + 1 < path_cost[next_y, next_x]
            ):
                path_cost[next_y, next_x] = current_path_cost + 1
                check_next.append((next_y, next_x))
    if path_cost[gameMapHeight - 1, gameMapWidth - 1] < 0:
        return False
    cur_y, cur_x = gameMapHeight - 1, gameMapWidth - 1
    next_val = path_cost[gameMapHeight - 1, gameMapWidth - 1] - 1
    reverse_path = []
    while True:
        if cur_x == 0 and cur_y == 0:
            break
        for dir_y, dir_x in dirs:
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if (
                next_y < 0
                or next_x < 0
                or next_x >= gameMapWidth
                or next_y >= gameMapHeight
                or gameMap[next_y, next_x] == CELL_CORRUPTED
            ):
                continue
            if path_cost[next_y, next_x] == next_val:
                reverse_path.append((next_y, next_x))
                next_val -= 1
                cur_x = next_x
                cur_y = next_y
    return reverse_path

def Part2(input: typing.TextIO, size: int = 71, steps: int = 1024) -> int:
    gameMap, data = loadField(input, size, steps)
    last_path = shortestPath(gameMap)

    for i in range(steps, len(data)):
        y, x = data[i]
        gameMap[y, x] = CELL_CORRUPTED
        if not (data[i] in last_path):
            continue
        path_cost = shortestPath(gameMap)
        if path_cost == False:
            break
        last_path = path_cost

    return data[i][::-1]
