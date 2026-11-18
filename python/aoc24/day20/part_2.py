import typing

from .shared import count_cheats


def Part2(input: typing.TextIO) -> int:
    return count_cheats(input, 100, 20)
