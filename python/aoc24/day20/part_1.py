import typing
from .shared import parseField, compute_shortest_path, count_cheats


def handle(input: typing.TextIO, cheatSavings: int) -> int:
    field = parseField(input)
    if field is None or field.start is None or field.end is None:
        raise ValueError("Error parsing field")
    path = compute_shortest_path(field)
    return count_cheats(path, cheatSavings)


def Part1(input: typing.TextIO) -> int:
    return handle(input, 100)
