import sys
import typing

from .shared import _DIR_DOWN, _DIR_LEFT, _DIR_RIGHT, _DIR_UP, parse_input, simulate


def Part2(input: typing.TextIO) -> int:
    field, width, height = parse_input(input)
    best = 0

    for x in range(width):
        best = max(best, simulate(field, width, height, complex(x, 0), _DIR_DOWN))
        best = max(
            best, simulate(field, width, height, complex(x, height - 1), _DIR_UP)
        )

    for y in range(height):
        best = max(best, simulate(field, width, height, complex(0, y), _DIR_RIGHT))
        best = max(
            best, simulate(field, width, height, complex(width - 1, y), _DIR_LEFT)
        )

    return best


if __name__ == "__main__":
    print(Part2(sys.stdin))
