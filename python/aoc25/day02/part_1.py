import sys
import typing
from .parse_input import _parse_input
from .find_repeated_blocks import _find_repeated_blocks


def Part1(input: typing.TextIO) -> int:
    total_sum = 0
    for start, end in _parse_input(input):
        invalids = _find_repeated_blocks(start, end)
        for num, k in invalids.items():
            if k == 2:
                total_sum += num
    return total_sum


if __name__ == "__main__":
    print("Part 1:", Part1(sys.stdin))
