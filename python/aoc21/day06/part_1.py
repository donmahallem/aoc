import sys
import typing

from shared import calculate_fish_counts, parse_input


def Part1(input_data: typing.TextIO) -> int:
    parsed_input = parse_input(input_data)
    return calculate_fish_counts(parsed_input, 80)


if __name__ == "__main__":
    print(Part1(sys.stdin))
