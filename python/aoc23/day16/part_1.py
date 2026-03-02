import sys
import typing
from .shared import parse_input, simulate, _DIR_RIGHT


def Part1(input: typing.TextIO) -> int:
    field, width, height = parse_input(input)
    return simulate(field, width, height, complex(0, 0), _DIR_RIGHT)


if __name__ == "__main__":
    print(Part1(sys.stdin))
