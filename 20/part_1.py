import codecs
import numpy as np

np.set_printoptions(linewidth=200)
test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = [line.strip() for line in f.readlines()]

field = np.zeros((len(data), len(data[0])))
WIDTH = field.shape[1]
HEIGHT = field.shape[1]
CELL_WALL = -1
CELL_START = -2
CELL_EMPTY = 0
player_position = None
end_position = None
for row in range(0, len(data)):
    for col in range(0, len(data[row])):
        if data[row][col] == "#":
            field[row, col] = CELL_WALL
        elif data[row][col] == "E":
            end_position = (row, col)
        elif data[row][col] == "S":
            player_position = (row, col)


def calculatePathCost(test_map):
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


def shortestPath(cost_map):
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    check_next = [end_position]
    cells = list([end_position])
    while len(check_next) > 0:
        cur_pos = check_next.pop(0)
        cur_y, cur_x = cur_pos
        current_value = cost_map[cur_y, cur_x]
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


path_cost = calculatePathCost(field)
normal_path_cost = path_cost[end_position[0], end_position[1]]
normal_path_taken = shortestPath(path_cost)

from tqdm import tqdm

summe = 0
savings = dict()
cheats = set()
for step_a in tqdm(range(0, len(normal_path_taken) - 1)):
    for step_b in range(len(normal_path_taken) - 1, step_a, -1):
        p1_y, p1_x = normal_path_taken[step_a]
        p2_y, p2_x = normal_path_taken[step_b]

        dst = abs(p1_x - p2_x) + abs(p1_y - p2_y)
        if dst != 2:
            continue
        saved = step_b - step_a - dst
        if path_cost[p1_y, p1_y] - path_cost[p1_y, p1_y] == saved:
            continue
        elif saved < 100:
            continue
        summe += 1
print(summe)
