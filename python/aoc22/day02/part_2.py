import sys
import typing

from .shared import _parser


def Part2(input_data: typing.TextIO) -> int:
    assert input_data is not None, "Input data must not be None"

    score: int = 0
    for l, r in _parser(input_data):
        score += (l + r + 2) % 3 + 1 + (r * 3)
    return score


if __name__ == "__main__":
    print(Part2(sys.stdin))
