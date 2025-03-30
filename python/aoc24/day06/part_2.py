import typing
import numpy as np

from  .part_1 import turnRight,parseField

def moveNext(field, player_position, obstacle):
    p_y, p_x, (test_dir_y, test_dir_x) = player_position
    while True:
        next_p_x, next_p_y = p_x + test_dir_x, p_y + test_dir_y
        if next_p_x < 0 or next_p_x >= field.shape[1] or next_p_y < 0 or next_p_y >= field.shape[0]:
            return False
        if field[next_p_y, next_p_x] == 1:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        elif next_p_y == obstacle[0] and next_p_x == obstacle[1]:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        else:
            break
    player_position = (next_p_y, next_p_x, (test_dir_y, test_dir_x))
    return player_position

def Part2(input: typing.TextIO) -> int:
    player_map,initial_player_position=parseField(input)

    circular_maps_num = 0
    for y in range(player_map.shape[0]):
            for x in range(player_map.shape[1]):
                player_position = initial_player_position
                if player_map[y, x] == 1 or (
                    player_position[0] == y and player_position[1] == x
                ):
                    continue
                path = set([player_position])
                while True:
                    next_pos = moveNext(player_map, player_position, (y, x))
                    if next_pos == False:
                        # outside play area
                        break
                    player_position = next_pos
                    if player_position in path:
                        circular_maps_num += 1
                        break
                    path.add(player_position)
    return circular_maps_num
