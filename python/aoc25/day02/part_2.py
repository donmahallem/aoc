import sys
import typing
from .parse_input import parseInputGen
from .find_repeated_blocks import find_repeated_blocks


def Part2(input: typing.TextIO) -> int:
    total_sum = 0
    for start, end in parseInputGen(input):
        invalids = find_repeated_blocks(start, end)
        # Sum all invalid numbers (with k >= 2)
        for num in invalids:
            total_sum += num
    return total_sum


if __name__ == "__main__":
    print(Part2(sys.stdin))
