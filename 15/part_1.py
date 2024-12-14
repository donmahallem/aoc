import codecs
import numpy as np

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.read()

data = data.split("\r\n\r\n")
data[0] = data[0].split("\r\n")
field = np.zeros((len(data[0]), len(data[0][0])))
CELL_WALL = 2
CELL_BOX = 1
CELL_EMPTY = 0
player_position = None
for row in range(0, len(data[0])):
    for col in range(0, len(data[0][row])):
        if data[0][row][col] == "#":
            field[row, col] = CELL_WALL
        elif data[0][row][col] == "O":
            field[row, col] = CELL_BOX
        elif data[0][row][col] == "@":
            player_position = (row, col)

moves = []
data[1] = "".join(data[1].split())


def map_movements(a):
    if a == "<":
        return (0, -1)
    elif a == ">":
        return (0, 1)
    elif a == "v":
        return (1, 0)
    elif a == "^":
        return (-1, 0)


movements = list(map(map_movements, data[1]))
print(player_position)


def next_empty(cur_y, cur_x, dir_y, dir_x):
    next_y, next_x = cur_y + dir_y, cur_x + dir_x
    if field[next_y, next_x] == CELL_WALL:
        return None
    elif field[next_y, next_x] == CELL_EMPTY:
        return (next_y, next_x)
    else:
        return next_empty(next_y, next_x, dir_y, dir_x)


for move in movements:
    next_y, next_x = player_position[0] + move[0], player_position[1] + move[1]
    if field[next_y, next_x] == CELL_WALL:
        continue
    elif field[next_y, next_x] == CELL_EMPTY:
        player_position = (next_y, next_x)
        continue
    elif field[next_y, next_x] == CELL_BOX:
        maybe_next_empty = next_empty(next_y, next_x, move[0], move[1])
        if maybe_next_empty:
            next_empty_y, next_empty_x = maybe_next_empty
            player_position = (next_y, next_x)
            field[next_y, next_x] = CELL_EMPTY
            field[next_empty_y, next_empty_x] = CELL_BOX
        else:
            continue

print(field)
boxes = list(zip(*np.where(field == 1)))
print(sum([y * 100 + x for y, x in boxes]))
