import typing
import numpy as np

CELL_WALL = -1
CELL_START = -2
CELL_EMPTY = 0

Position: typing.TypeAlias = tuple[int, int]


NpData = np.ndarray[tuple[int, int], np.dtype[np.uint8]]


def parseField(input: typing.TextIO) -> tuple[NpData, Position | None, Position | None]:
    data = [line.strip() for line in input.readlines()]
    field: NpData = typing.cast(NpData, np.zeros((len(data), len(data[0]))))
    player_position: Position | None = None
    end_position: Position | None = None
    for row in range(0, len(data)):
        for col in range(0, len(data[row])):
            if data[row][col] == "#":
                field[row, col] = CELL_WALL
            elif data[row][col] == "E":
                end_position = (row, col)
            elif data[row][col] == "S":
                player_position = (row, col)
    return field, player_position, end_position


def calculatePathCost(
    test_map: np.typing.NDArray, player_position: Position
) -> np.typing.NDArray:
    check_next = [(0, player_position)]
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    path_cost = np.zeros(test_map.shape, dtype=np.int32) - 1
    path_cost[player_position[0], player_position[1]] = 0
    while len(check_next) > 0:
        last_val, cur_pos = check_next.pop(0)
        cur_y, cur_x = cur_pos
        for dir in dirs:
            dir_y, dir_x = dir
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if test_map[next_y, next_x] == CELL_WALL:
                continue
            next_path_value = last_val + 1
            if (
                path_cost[next_y, next_x] < 0
                or path_cost[next_y, next_x] > next_path_value
            ):
                path_cost[next_y, next_x] = next_path_value
                check_next.append((next_path_value, (next_y, next_x)))
    return path_cost


def shortestPath(field, path_cost, end_position):
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    check_next = [end_position]
    cells = list([end_position])
    while len(check_next) > 0:
        cur_pos = check_next.pop(0)
        cur_y, cur_x = cur_pos
        current_value = path_cost[cur_y, cur_x]
        for dir in dirs:
            dir_y, dir_x = dir
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if field[next_y, next_x] == CELL_WALL or (next_y, next_x) in cells:
                continue
            if (
                path_cost[next_y, next_x] == current_value - 1
                or path_cost[next_y, next_x] == current_value - 1001
            ):
                check_next.append((next_y, next_x))
                cells.append((next_y, next_x))

    return cells


def CountSheats(normal_path_taken, path_cost, cheat_savings: int) -> int:
    summe = 0
    for step_a in range(0, len(normal_path_taken) - 1):
        for step_b in range(len(normal_path_taken) - 1, step_a, -1):
            p1_y, p1_x = normal_path_taken[step_a]
            p2_y, p2_x = normal_path_taken[step_b]

            dst = abs(p1_x - p2_x) + abs(p1_y - p2_y)
            if dst != 2:
                continue
            saved = step_b - step_a - dst
            if path_cost[p1_y, p1_y] - path_cost[p1_y, p1_y] == saved:
                continue
            elif saved < cheat_savings:
                continue
            summe += 1
    return summe
