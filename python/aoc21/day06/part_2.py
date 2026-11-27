import sys
import typing

from .shared import calculate_fish_counts, parse_input


def Part2(input_data: typing.TextIO) -> int:
    parsed_input = parse_input(input_data)
    return calculate_fish_counts(parsed_input, 256)


if __name__ == "__main__":
    print(Part2(sys.stdin))
