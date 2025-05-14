import typing
from .part_1 import loadField, calculatePathCost


def getInverseSamplePath(cost_map_dict, target_value, end_position):
    dirs = (0, 1), (1, 0), (0, -1), (-1, 0)
    check_next = [(target_value, None, end_position)]
    cells = list([end_position])
    while len(check_next) > 0:
        current_value, cur_dir, cur_pos = check_next.pop(0)
        cur_y, cur_x = cur_pos
        for dir in dirs:
            dir_y, dir_x = dir
            next_y, next_x = cur_y + dir_y, cur_x + dir_x
            next_pos = (next_y, next_x)
            if next_pos not in cost_map_dict or (next_y, next_x) in cells:
                continue
            # if last_val and last_dir==dir:
            if (current_value - 1) in cost_map_dict[next_pos]:
                check_next.append((current_value - 1, dir, (next_y, next_x)))
                cells.append((next_y, next_x))
            elif (current_value - 1001) in cost_map_dict[next_pos]:
                check_next.append(
                    (current_value - 1001, dir, (next_y, next_x)))
                cells.append((next_y, next_x))
    return cells


def Part2(input: typing.TextIO) -> int:
    field, start, end = loadField(input)

    path_cost_dict = calculatePathCost(field, start, end)
    shortest_path_cost = min(path_cost_dict[end[0], end[1]])

    turns = shortest_path_cost // 1000
    straights = (shortest_path_cost - turns) % 1000
    inverse = getInverseSamplePath(path_cost_dict, shortest_path_cost, end)
    return len(inverse)
