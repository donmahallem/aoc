import typing
from .shared import _parse_field, _compute_shortest_path, _count_cheats


def handle(input: typing.TextIO, cheatSavings: int) -> int:
    field = _parse_field(input)
    if field is None or field.start is None or field.end is None:
        raise ValueError("Error parsing field")
    path = _compute_shortest_path(field)
    return _count_cheats(path, cheatSavings)


def Part1(input: typing.TextIO) -> int:
    return handle(input, 100)
