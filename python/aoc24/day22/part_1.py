import typing
from .shared import _parseInput, _step, _ITERATIONS


def Part1(input: typing.TextIO) -> int:
    data = _parseInput(input)
    total = 0
    for seed in data:
        tmp = seed
        # manually unroll loop in C‑style for speed
        for _ in range(_ITERATIONS):
            tmp = _step(tmp)
        total += tmp
    return total
