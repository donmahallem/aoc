import sys
import typing
from .shared import parse_input, find_shortest_path


def Part2(input: typing.TextIO) -> int:
    cells, width, height = parse_input(input)
    return find_shortest_path(cells, width, height, 4, 10)


if __name__ == "__main__":
    print(Part2(sys.stdin))
