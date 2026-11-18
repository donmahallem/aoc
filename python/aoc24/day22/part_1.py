import typing

from .shared import _loop_steps, _parseInput


def Part1(input: typing.TextIO) -> int:
    data = _parseInput(input)
    total = 0
    for seed in data:
        tmp = _loop_steps(seed)
        total += tmp
    return total
