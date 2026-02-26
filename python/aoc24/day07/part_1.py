import typing
from .shared import parseRows, Equation


def is_valid(target: int, values: tuple[int, ...]) -> bool:
    if len(values) == 1:
        return values[0] == target

    last = values[-1]
    remaining = values[:-1]

    if target > last:
        if is_valid(target - last, remaining):
            return True

    if target % last == 0:
        if is_valid(target // last, remaining):
            return True

    return False


def Part1(file_stream: typing.TextIO) -> int:
    rows = parseRows(file_stream)
    total_sum = 0

    for expected, values in rows:
        if is_valid(expected, values):
            total_sum += expected

    return total_sum
