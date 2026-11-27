import sys
import typing

from .shared import decode_mapping, parse_input


def Part1(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    count = 0
    for line in data:
        mapping_decoded = decode_mapping(line.mapping)
        for item in line.output:
            if mapping_decoded[item] in (1, 4, 7, 8):
                count += 1
    return count


if __name__ == "__main__":
    print(Part1(sys.stdin))
