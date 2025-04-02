import typing
from .part_1 import parseInput


def walk(data_np, y, x, looking_for, ends):
    sum = 0
    dirs = [(0, 1), (1, 0), (-1, 0), (0, -1)]
    for dir_y, dir_x in dirs:
        check_x, check_y = dir_x + x, dir_y + y
        if (
            check_x < 0
            or check_y < 0
            or check_y >= data_np.shape[0]
            or check_x >= data_np.shape[1]
        ):
            continue
        if looking_for == 9 and data_np[check_y, check_x] == 9:
            sum += 1
            ends.add((check_y, check_x))
        elif data_np[check_y, check_x] == looking_for:
            sum += walk(data_np, check_y, check_x, looking_for + 1, ends)
    return sum


def Part2(input: typing.TextIO) -> int:
    data_np, trailheads = parseInput(input)
    summe = 0
    for trailhead in trailheads:
        k = set()
        summe += walk(data_np, trailhead[0], trailhead[1], 1, k)
        # summe+=len(k)
    return summe
