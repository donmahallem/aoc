import typing
import numpy as np
from .shared import CELL_BOX, CELL_WALL, CELL_EMPTY, translateMovement


def parseField(input: typing.TextIO):
    data = list(map(str.strip, input.readlines()))
    splitIdx = data.index("")
    fieldData = data[0:splitIdx]
    movementData = data[splitIdx + 1 :]
    field = np.zeros((len(fieldData), len(fieldData[0])))
    player_position = None
    for row in range(0, len(fieldData)):
        for col in range(0, len(fieldData[row])):
            if fieldData[row][col] == "#":
                field[row, col] = CELL_WALL
            elif fieldData[row][col] == "O":
                field[row, col] = CELL_BOX
            elif fieldData[row][col] == "@":
                player_position = (row, col)

    movementDataStr = "".join(movementData)

    return field, player_position, list(map(translateMovement, movementDataStr))


def next_empty(field, cur_y, cur_x, dir_y, dir_x):
    next_y, next_x = cur_y + dir_y, cur_x + dir_x
    if field[next_y, next_x] == CELL_WALL:
        return None
    elif field[next_y, next_x] == CELL_EMPTY:
        return (next_y, next_x)
    else:
        return next_empty(field, next_y, next_x, dir_y, dir_x)


def Part1(input: typing.TextIO) -> int:
    field, playerPosition, movements = parseField(input)

    for move in movements:
        next_y, next_x = playerPosition[0] + move[0], playerPosition[1] + move[1]
        if field[next_y, next_x] == CELL_WALL:
            continue
        elif field[next_y, next_x] == CELL_EMPTY:
            playerPosition = (next_y, next_x)
            continue
        elif field[next_y, next_x] == CELL_BOX:
            maybe_next_empty = next_empty(field, next_y, next_x, move[0], move[1])
            if maybe_next_empty:
                next_empty_y, next_empty_x = maybe_next_empty
                playerPosition = (next_y, next_x)
                field[next_y, next_x] = CELL_EMPTY
                field[next_empty_y, next_empty_x] = CELL_BOX
            else:
                continue

    boxes = list(zip(*np.where(field == 1)))
    return sum([y * 100 + x for y, x in boxes])
