import typing
import numpy as np

from .part_1 import turnRight, parseField


def moveNext(field, player_position, obstacle: tuple[int, int] | None):
    p_y, p_x, (test_dir_y, test_dir_x) = player_position
    while True:
        next_p_x, next_p_y = p_x + test_dir_x, p_y + test_dir_y
        if (next_p_x < 0 or next_p_x >= field.shape[1] or next_p_y < 0
                or next_p_y >= field.shape[0]):
            return False
        if field[next_p_y, next_p_x] == 1:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        elif obstacle is not None and next_p_y == obstacle[0] and next_p_x == obstacle[1]:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        else:
            break
    player_position = (next_p_y, next_p_x, (test_dir_y, test_dir_x))
    return player_position


def simulate_with_obstacle(field, start_pos, obstacle: tuple[int, int]) -> bool:
    seen = set()
    pos = start_pos
    while True:
        if pos in seen:
            return True  # loop detected
        seen.add(pos)
        next_pos = moveNext(field, pos, obstacle)
        if next_pos is False:
            return False
        pos = next_pos


def Part2(input: typing.TextIO) -> int:
    player_map, initial_player_position = parseField(input)
    if initial_player_position is None:
        raise Exception("No guard found")

    # First, trace the guard's path without extra obstacles to limit candidates.
    base_positions: list[tuple[int, int]] = []
    pos = initial_player_position
    while True:
        next_pos = moveNext(player_map, pos, None)
        if next_pos is False:
            break
        pos = next_pos
        base_positions.append((pos[0], pos[1]))

    candidates = set(base_positions)
    # Do not place an obstacle on the starting cell.
    candidates.discard((initial_player_position[0], initial_player_position[1]))

    circular_maps_num = 0
    for obstacle in candidates:
        if simulate_with_obstacle(player_map, initial_player_position, obstacle):
            circular_maps_num += 1

    return circular_maps_num
