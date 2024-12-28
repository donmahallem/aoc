import codecs
import numpy as np

np.set_printoptions(linewidth=200)

CELL_WALL = -1
CELL_START = -2
CELL_EMPTY = 0


def loadField(filepath):
    with codecs.open(filepath) as f:
        data = [line.strip() for line in f.readlines()]
    field = np.zeros((len(data), len(data[0])))
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
    return field, player_position, end_position


def calculatePathCost(field, start_pos, end_pos):
    check_next = [(0, (0, 1), [start_pos])]
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    cost_dict = dict()
    cost_dict[start_pos] = set([0])
    while len(check_next) > 0:
        current_cost, last_dir, cur_path = check_next.pop(0)
        cur_pos = cur_path[-1]
        if cur_pos == end_pos:
            continue

        cur_y, cur_x = cur_pos

        for dir in dirs:
            dir_y, dir_x = dir
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            if field[next_y, next_x] == CELL_WALL:
                continue
            if (next_y, next_x) in cur_path:
                continue
            next_cost = current_cost + (1001 if last_dir != dir else 1)
            next_pos = (next_y, next_x)
            if next_pos in cost_dict and min(cost_dict[next_pos]) + 1001 < next_cost:
                continue
            if next_pos in cost_dict:
                cost_dict[next_pos].add(next_cost)
            else:
                cost_dict[next_pos] = set([next_cost])
            check_next.append((next_cost, dir, cur_path + [next_pos]))
    return cost_dict


if __name__ == "__main__":
    test_data = False
    field, start, end = loadField("data.txt" if test_data else "data2.txt")

    path_cost_map = calculatePathCost(field, start, end)
    print(min(path_cost_map[end]))
