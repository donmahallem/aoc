import sys
import typing

from .shared import parse_ast

def Part1(input_data: typing.TextIO) -> int:
    count = 0
    for line in input_data:
        line = line.strip()
        # Process each line as needed
        result, _ = parse_ast(line)
        count += result

    return count


if __name__ == "__main__":
    print(Part1(sys.stdin))
