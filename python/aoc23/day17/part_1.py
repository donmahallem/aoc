import sys
import typing
from .shared import _parse_input, _find_shortest_path


def Part1(input: typing.TextIO) -> int:
    cells, width, height = _parse_input(input)
    return _find_shortest_path(cells, width, height, 0, 3)


if __name__ == "__main__":
    print(Part1(sys.stdin))
