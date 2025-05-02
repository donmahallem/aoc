import typing
import numpy as np
from .shared import (
    CELL_EMPTY,
    CELL_WALL,
    CELL_BOX_LEFT,
    CELL_BOX_RIGHT,
    translateMovement,
)


def parseField(input: typing.TextIO):
    inputData = list(map(str.strip, input.readlines()))
    splitIdx = inputData.index("")
    fieldData = inputData[0:splitIdx]
    movementData = inputData[splitIdx + 1 :]
    field = np.zeros((len(fieldData), len(fieldData[0] * 2)))
    player_position = None
    for row in range(0, len(fieldData)):
        for col in range(0, len(fieldData[row])):
            if fieldData[row][col] == "#":
                field[row, 2 * col : 2 * col + 2] = [CELL_WALL, CELL_WALL]
            elif fieldData[row][col] == "O":
                field[row, 2 * col : 2 * col + 2] = [CELL_BOX_LEFT, CELL_BOX_RIGHT]
            elif fieldData[row][col] == "@":
                player_position = (row, col * 2)

    movementDataStr = "".join(movementData)
    return field, player_position, list(map(translateMovement, movementDataStr))


def next_empty(
    field: np.typing.NDArray, cur_y: int, cur_x: int, dir_y: int, dir_x: int
) -> tuple[int, int] | None:
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
        return next_empty(field, next_y, next_x, dir_y, dir_x)


def getMoveableY(field, cur_pos, dir):
    dir_y = dir[0]
    cur_y, cur_x = cur_pos
    corrected_x = cur_x
    if field[cur_y, cur_x] == CELL_BOX_RIGHT:
        corrected_x -= 1
    if field[cur_y, corrected_x] == CELL_WALL:
        return None
    elif field[cur_y, corrected_x] == CELL_EMPTY:
        return []
    next_y = cur_y + dir_y
    items_to_move = [(cur_y, corrected_x)]
    if field[next_y, corrected_x] == CELL_BOX_LEFT:
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


def moveY(field, cur_pos, dir) -> bool:
    next_pos = cur_pos[0] + dir[0], cur_pos[1] + dir[1]
    items = getMoveableY(field, next_pos, dir)
    if items == None or len(items) == 0:
        return False
    # print("Move", sorted(items, key=lambda a: a[0], reverse=dir[0] > 0))
    for y, x in sorted(items, key=lambda a: a[0], reverse=dir[0] > 0):
        field[y + dir[0], x] = CELL_BOX_LEFT
        field[y + dir[0], x + 1] = CELL_BOX_RIGHT
        field[y, x : x + 2] = CELL_EMPTY
    return True


def printField(field, playerPosition):
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
            elif playerPosition == (y, x):
                line += "@"
            else:
                line += "-"
        lines.append(line)
    print("\r\n".join(lines))


def Part2(input: typing.TextIO) -> int:
    field, playerPosition, movements = parseField(input)
    for move in movements:
        next_y, next_x = playerPosition[0] + move[0], playerPosition[1] + move[1]
        if field[next_y, next_x] == CELL_WALL:
            continue
        elif field[next_y, next_x] == CELL_EMPTY:
            playerPosition = (next_y, next_x)
            continue
        elif (
            field[next_y, next_x] == CELL_BOX_LEFT
            or field[next_y, next_x] == CELL_BOX_RIGHT
        ):
            if move[1] != 0:
                maybe_next_empty = next_empty(field, next_y, next_x, move[0], move[1])
                if maybe_next_empty:
                    next_empty_y, next_empty_x = maybe_next_empty
                    playerPosition = (next_y, next_x)
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
                if moveY(field, playerPosition, move):
                    playerPosition = (next_y, next_x)
            else:
                continue
    boxes = list(zip(*np.where(field == CELL_BOX_LEFT)))
    return sum([y * 100 + x for y, x in boxes])
