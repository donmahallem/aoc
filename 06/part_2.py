import codecs
import numpy as np
from tqdm import tqdm

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

rows = len(data)
columns = len(data)

player_map = np.zeros((rows, columns), dtype=np.uint8)
initial_player_position = None
for y in range(0, rows):
    for x in range(0, columns):
        if data[y][x] == "#":
            player_map[y, x] = 1
        elif data[y][x] == "^":
            initial_player_position = (y, x, (-1, 0))
        elif data[y][x] == ">":
            initial_player_position = (y, x, (0, 1))
        elif data[y][x] == "v":
            initial_player_position = (y, x, (-1, 0))
        elif data[y][x] == "<":
            initial_player_position = (y, x, (0, -1))
# print("Initial player position",initial_player_position)


def turnRight(cur_y, cur_x):
    if cur_x == 0 and cur_y == 1:
        return (0, -1)
    elif cur_x == 1 and cur_y == 0:
        return (1, 0)
    elif cur_x == 0 and cur_y == -1:
        return (0, 1)
    elif cur_x == -1 and cur_y == 0:
        return (-1, 0)
    else:
        raise IndexError(f"Invalid dir {cur_y},{cur_x}")


def moveNext(field, player_position, obstacle):
    p_y, p_x, (test_dir_y, test_dir_x) = player_position
    while True:
        next_p_x, next_p_y = p_x + test_dir_x, p_y + test_dir_y
        if next_p_x < 0 or next_p_x >= columns or next_p_y < 0 or next_p_y >= rows:
            return False
        if field[next_p_y, next_p_x] == 1:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        elif next_p_y == obstacle[0] and next_p_x == obstacle[1]:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        else:
            break
    player_position = (next_p_y, next_p_x, (test_dir_y, test_dir_x))
    return player_position


circular_maps_num = 0
with tqdm(total=rows * columns) as pbar:
    for y in range(rows):
        for x in range(columns):
            player_position = initial_player_position
            if player_map[y, x] == 1 or (player_position[0]==y and player_position[1]==x):
                pbar.update(1)
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
            pbar.update(1)
print(circular_maps_num)
