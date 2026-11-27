import sys
import typing

from .shared import decode_mapping, parse_input


def Part2(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    total = 0
    for line in data:
        mapping_decoded = decode_mapping(line.mapping)
        value = 0
        for item in line.output:
            value = value * 10 + mapping_decoded[item]
        total += value
    return total


if __name__ == "__main__":
    print(Part2(sys.stdin))
