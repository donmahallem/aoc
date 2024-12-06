import codecs
import numpy as np

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

rows = len(data)
columns = len(data)

player_map = np.zeros((rows, columns), dtype=np.uint8)
player_position = None
for y in range(0, rows):
    for x in range(0, columns):
        if data[y][x] == "#":
            player_map[y, x] = 1
        elif data[y][x] == "^":
            player_position = (y, x, (-1, 0))
        elif data[y][x] == ">":
            player_position = (y, x, (0, 1))
        elif data[y][x] == "v":
            player_position = (y, x, (-1, 0))
        elif data[y][x] == "<":
            player_position = (y, x, (0, -1))


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


def moveNext(field, player_position):
    p_y, p_x, (test_dir_y, test_dir_x) = player_position
    while True:
        next_p_x, next_p_y = p_x + test_dir_x, p_y + test_dir_y
        if next_p_x < 0 or next_p_x >= columns or next_p_y < 0 or next_p_y >= rows:
            return False
        if field[next_p_y, next_p_x] == 1:
            # print("turn right")
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        else:
            break
    player_position = (next_p_y, next_p_x, (test_dir_y, test_dir_x))
    return player_position


stepper = 0

path = set([(player_position[0], player_position[1])])
while True:
    stepper += 1
    next_pos = moveNext(player_map, player_position)
    if next_pos == False:
        break
    player_position = next_pos
    p_y, p_x, _ = player_position
    path.add((p_y, p_x))
    if p_x < 0 or p_y < 0 or p_x >= columns or p_y >= rows:
        break
print(len(path))
