import sys
import typing

from .shared import _parser


def Part1(input_data: typing.TextIO) -> int:
    assert input_data is not None, "Input data must not be None"

    score: int = 0
    for l, r in _parser(input_data):
        score += r + 1
        if l == r:
            score += 3
        elif (l + 1) % 3 == r:
            score += 6
    return score


if __name__ == "__main__":
    print(Part1(sys.stdin))
