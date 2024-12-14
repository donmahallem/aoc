import codecs
import numpy as np

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.read()

data = data.split("\r\n\r\n")
data[0] = data[0].split("\r\n")
field = np.zeros((len(data[0]), len(data[0][0]) * 2))
CELL_WALL = 2
CELL_BOX_LEFT = 1
CELL_BOX_RIGHT = 3
CELL_EMPTY = 0
player_position = None
for row in range(0, len(data[0])):
    for col in range(0, len(data[0][row])):
        if data[0][row][col] == "#":
            field[row, 2 * col : 2 * col + 2] = [CELL_WALL, CELL_WALL]
        elif data[0][row][col] == "O":
            field[row, 2 * col : 2 * col + 2] = [CELL_BOX_LEFT, CELL_BOX_RIGHT]
        elif data[0][row][col] == "@":
            player_position = (row, col * 2)

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


print(field)
player_movements = list(map(map_movements, data[1]))
print(player_position)


def next_empty(cur_y, cur_x, dir_y, dir_x):
    next_y, next_x = cur_y + dir_y, cur_x + dir_x
    if field[next_y, next_x] == CELL_WALL:
        return None
    elif dir_y != 0 and (
        (
            field[next_y, next_x] == CELL_BOX_LEFT
            and field[cur_y, cur_x] == CELL_BOX_RIGHT
        )
        or (
            field[next_y, next_x] == CELL_BOX_RIGHT
            and field[cur_y, cur_x] == CELL_BOX_LEFT
        )
    ):
        return None
    elif field[next_y, next_x] == CELL_EMPTY:
        if dir_y != 0:
            if (
                field[cur_y, cur_x] == CELL_BOX_LEFT
                and field[next_y, next_x + 1] != CELL_EMPTY
            ):
                return None
            elif (
                field[cur_y, cur_x] == CELL_BOX_RIGHT
                and field[next_y, next_x - 1] != CELL_EMPTY
            ):
                return None
        return (next_y, next_x)
    else:
        return next_empty(next_y, next_x, dir_y, dir_x)


def isMovePossibleY(cur_y, cur_x, dir_y):
    if field[cur_y, cur_x] == CELL_BOX_RIGHT:
        return isMovePossibleY(cur_y, cur_x - 1, dir_y)
    next_y, next_x1, next_x2 = cur_y + dir_y, cur_x, cur_x + 1
    if field[next_y, next_x1] == CELL_WALL or field[next_y, next_x2] == CELL_WALL:
        return False
    elif field[next_y, next_x1] == CELL_EMPTY and field[next_y, next_x2] == CELL_EMPTY:
        return True
    else:
        return isMovePossibleY(next_y, next_x1, dir_y) and isMovePossibleY(
            next_y, next_x2, dir_y
        )


def moveYTree(cur_y, cur_x, dir_y):
    if field[cur_y, cur_x] == CELL_BOX_RIGHT:
        moveYTree(cur_y, cur_x - 1, dir_y)
        return
    next_y, next_x1, next_x2 = cur_y + dir_y, cur_x, cur_x + 1
    if field[next_y, next_x1] == CELL_BOX_LEFT:
        moveYTree(next_y, next_x1, dir_y)
    else:
        if field[next_y, next_x1] == CELL_BOX_RIGHT:
            moveYTree(next_y, next_x1 - 1, dir_y)
        if field[next_y, next_x2] == CELL_BOX_LEFT:
            moveYTree(next_y, next_x2, dir_y)
    field[next_y, next_x1] = CELL_BOX_LEFT
    field[next_y, next_x2] = CELL_BOX_RIGHT
    field[cur_y, cur_x : cur_x + 2] = 0


def moveIfPossible(cur_y, cur_x, dir_y, dir_x, depth=0):
    if dir_x != 0 and dir_y != 0:
        raise "No diagonal moves"
    next_y, next_x = cur_y + dir_y, cur_x + dir_x
    if dir_x != 0:
        if field[next_y, next_x] == CELL_WALL:
            return False
        elif field[next_y, next_x] == CELL_EMPTY:
            field[next_y, next_x] = field[cur_y, cur_x]
            if depth == 0:
                field[cur_y, cur_x] = CELL_EMPTY
            return True
        elif (
            field[next_y, next_x] == CELL_BOX_LEFT
            or field[next_y, next_x] == CELL_BOX_RIGHT
        ) and moveIfPossible(next_y, next_x, dir_y, dir_x, depth + 1):
            if depth == 0:
                field[cur_y, cur_x] = CELL_EMPTY
            return True


def printField(field):
    lines = []
    for y in range(field.shape[0]):
        line = ""
        for x in range(field.shape[1]):
            if field[y, x] == CELL_BOX_LEFT:
                line += "["
            elif field[y, x] == CELL_BOX_RIGHT:
                line += "]"
            elif field[y, x] == CELL_WALL:
                line += "#"
            elif player_position == (y, x):
                line += "@"
            else:
                line += "-"
        lines.append(line)
    print("\r\n".join(lines))


def printDir(move):
    if move == (0, 1):
        print(">")
    elif move == (0, -1):
        print("<")
    elif move == (-1, 0):
        print("^")
    else:
        print("v")


def getMoveableY(field, cur_pos, dir):
    dir_y = dir[0]
    cur_y,cur_x=cur_pos
    corrected_x = cur_x
    if field[cur_y, cur_x] == CELL_BOX_RIGHT:
        corrected_x-=1
    if (
        field[cur_y, corrected_x] == CELL_WALL
    ):
        return None
    elif (
        field[cur_y, corrected_x] == CELL_EMPTY
    ):
        return []
    next_y = cur_y + dir_y
    items_to_move = [(cur_y, corrected_x)]
    if field[next_y,corrected_x]==CELL_BOX_LEFT:
            res = getMoveableY(field, (next_y, corrected_x), dir)
            if res == None:
                return None
            items_to_move.extend(res)
    else:
        for i in range(2):
            res = getMoveableY(field, (next_y, corrected_x + i), dir)
            if res == None:
                return None
            items_to_move.extend(res)
    return items_to_move


def moveY(field, cur_pos, dir):
    next_pos = cur_pos[0] + dir[0], cur_pos[1] + dir[1]
    items = getMoveableY(field, next_pos, dir)
    if items == None or len(items) == 0:
        return False
    #print("Move", sorted(items, key=lambda a: a[0], reverse=dir[0] > 0))
    for y, x in sorted(items, key=lambda a: a[0], reverse=dir[0] > 0):
        field[y + dir[0], x] = CELL_BOX_LEFT
        field[y + dir[0], x + 1] = CELL_BOX_RIGHT
        field[y, x : x + 2] = CELL_EMPTY
    return True


for move in player_movements:
    #printField(field)
    #input()
    #printDir(move)
    next_y, next_x = player_position[0] + move[0], player_position[1] + move[1]
    if field[next_y, next_x] == CELL_WALL:
        continue
    elif field[next_y, next_x] == CELL_EMPTY:
        player_position = (next_y, next_x)
        continue
    elif (
        field[next_y, next_x] == CELL_BOX_LEFT
        or field[next_y, next_x] == CELL_BOX_RIGHT
    ):
        if move[1] != 0:
            maybe_next_empty = next_empty(next_y, next_x, move[0], move[1])
            if maybe_next_empty:
                next_empty_y, next_empty_x = maybe_next_empty
                player_position = (next_y, next_x)
                # print("Move to:", next_empty_y, next_empty_x)
                if move[1] < 0:
                    field[next_y, next_empty_x:next_x] = field[
                        next_y, next_empty_x + 1 : next_x + 1
                    ]
                    field[next_y, next_x] = CELL_EMPTY
                elif move[1] > 0:
                    field[next_y, next_x + 1 : next_empty_x + 1] = field[
                        next_y, next_x:next_empty_x
                    ]
                    field[next_y, next_x] = CELL_EMPTY
        elif move[0] != 0:
            if moveY(field, player_position, move):
                player_position = (next_y, next_x)
        else:
            continue
boxes = list(zip(*np.where(field == 1)))
print(sum([y * 100 + x for y, x in boxes]))
