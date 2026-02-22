import typing
from collections import deque
from .shared import CELL_CORRUPTED, loadField


def shortestPath(gameMap: list[list[int]]):
    """Return reverse shortest path from bottom-right to start or False if blocked"""
    height = len(gameMap)
    width = len(gameMap[0]) if height else 0
    # cost grid
    path_cost = [[-1] * width for _ in range(height)]
    dq = deque()
    dq.append((0, 0))
    path_cost[0][0] = 0
    dirs = ((0, 1), (1, 0), (0, -1), (-1, 0))
    while dq:
        cur_y, cur_x = dq.popleft()
        if cur_y == height - 1 and cur_x == width - 1:
            break
        current_path_cost = path_cost[cur_y][cur_x]
        for dir_y, dir_x in dirs:
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if (0 <= next_y < height and 0 <= next_x < width
                    and gameMap[next_y][next_x] != CELL_CORRUPTED):
                prev = path_cost[next_y][next_x]
                if prev < 0 or current_path_cost + 1 < prev:
                    path_cost[next_y][next_x] = current_path_cost + 1
                    dq.append((next_y, next_x))
    goal_cost = path_cost[height - 1][width - 1]
    if goal_cost < 0:
        return False
    cur_y, cur_x = height - 1, width - 1
    next_val = goal_cost - 1
    reverse_path: list[tuple[int, int]] = []
    while not (cur_y == 0 and cur_x == 0):
        for dir_y, dir_x in dirs:
            ny, nx = cur_y + dir_y, cur_x + dir_x
            if (0 <= ny < height and 0 <= nx < width
                    and gameMap[ny][nx] != CELL_CORRUPTED
                    and path_cost[ny][nx] == next_val):
                reverse_path.append((ny, nx))
                cur_y, cur_x = ny, nx
                next_val -= 1
                break
    return reverse_path


def Part2(input: typing.TextIO,
          size: int = 71,
          steps: int = 1024) -> list[int]:
    gameMap, data = loadField(input, size, steps)
    last_path = shortestPath(gameMap)
    path_set = set(last_path) if last_path else set()

    for i in range(steps, len(data)):
        y, x = data[i]
        gameMap[y][x] = CELL_CORRUPTED
        if (y, x) not in path_set:
            # corrupted cell
            continue
        new_path = shortestPath(gameMap)
        if new_path is False:
            break
        last_path = new_path
        path_set = set(new_path)

    # return first failing coordinate converted back to (x,y)
    return list(data[i][::-1])
