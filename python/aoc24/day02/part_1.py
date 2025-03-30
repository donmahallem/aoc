import sys
import typing
from .shared import checkRow


def Part1(input: typing.TextIO) -> int:
    data = input.readlines()

    safe_count = 0
    for dataline in data:
        row_data = [int(d) for d in dataline.split()]
        if checkRow(row_data):
            safe_count += 1
    return safe_count


if __name__ == "__main__":
    Part1(sys.stdin)
