import typing
from collections import deque
from .shared import CELL_CORRUPTED, loadField


def shortestPath(gameMap: list[list[int]]) -> list[list[int]]:
    """Compute shortest distance from (0,0) to every reachable cell.

    Uses a simple BFS queue (deque) on a python list-of-lists map. Returns a
    2‑D list of costs (or -1 for unreachable)."""
    height = len(gameMap)
    width = len(gameMap[0]) if height else 0
    # initialize cost grid with -1
    pathCost = [[-1] * width for _ in range(height)]
    dq = deque()
    dq.append((0, 0))
    pathCost[0][0] = 0
    dirs = ((0, 1), (1, 0), (0, -1), (-1, 0))
    while dq:
        cur_y, cur_x = dq.popleft()
        # reachted bottom right; stop search since we have all costs <= goal_cost
        if cur_y == height - 1 and cur_x == width - 1:
            break
        current_path_cost = pathCost[cur_y][cur_x]
        for dir_y, dir_x in dirs:
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if (0 <= next_y < height and 0 <= next_x < width
                    and gameMap[next_y][next_x] != CELL_CORRUPTED):
                prev = pathCost[next_y][next_x]
                if prev < 0 or current_path_cost + 1 < prev:
                    pathCost[next_y][next_x] = current_path_cost + 1
                    dq.append((next_y, next_x))
    return pathCost


def Part1(input: typing.TextIO, size: int = 71, steps: int = 1024) -> int:
    gameMap, _ = loadField(input, size, steps)
    pathCost = shortestPath(gameMap)
    return pathCost[-1][-1]
