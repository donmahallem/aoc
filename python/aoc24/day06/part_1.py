import typing
import numpy as np

Guard = typing.Tuple[int, int, typing.Tuple[int, int]]


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


def moveNext(
    field: np.typing.NDArray, player_position: Guard
) -> Guard | typing.Literal[False]:
    p_y, p_x, (test_dir_y, test_dir_x) = player_position
    while True:
        next_p_x, next_p_y = p_x + test_dir_x, p_y + test_dir_y
        if (
            next_p_x < 0
            or next_p_x >= field.shape[1]
            or next_p_y < 0
            or next_p_y >= field.shape[0]
        ):
            return False
        if field[next_p_y, next_p_x] == 1:
            test_dir_y, test_dir_x = turnRight(test_dir_y, test_dir_x)
        else:
            break
    player_position = (next_p_y, next_p_x, (test_dir_y, test_dir_x))
    return player_position


InputField = np.ndarray[tuple[int, int], np.dtype[np.uint8]]


def parseField(input: typing.TextIO) -> tuple[InputField, Guard | None]:
    data = [a.strip() for a in input.readlines()]

    rows = len(data)
    columns = len(data[0])
    player_map: InputField = np.zeros((rows, columns), dtype=np.uint8)
    initial_player_position: Guard | None = None
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
    return player_map, initial_player_position


def Part1(input: typing.TextIO) -> int:
    player_map, player_position = parseField(input)

    if player_position == None:
        raise Exception("No guard found")
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
        if (
            p_x < 0
            or p_y < 0
            or p_x >= player_map.shape[0]
            or p_y >= player_map.shape[1]
        ):
            break
    return len(path)
