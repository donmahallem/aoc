import sys
import typing

from .shared import parse_input


def Part1(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    median = sorted(data)[len(data) // 2]
    return sum(abs(x - median) for x in data)


if __name__ == "__main__":
    print(Part1(sys.stdin))
