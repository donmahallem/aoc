import sys
import typing

from .shared import row_generator


def reduce(a: list[int]) -> int:
    base_val = 0
    base_result = 0
    base_val2 = 0
    for idx in range(len(a) - 1):
        ch = a[idx]
        if base_val < ch:
            base_val = ch
            base_result = ch * 10
            base_val2 = 0
        for idx2 in range(idx + 1, len(a)):
            base_val2 = max(base_val2, a[idx2])
    return base_result + base_val2


def Part1(input: typing.TextIO) -> int:
    total_sum = 0
    for line in row_generator(input):
        total_sum += reduce(line)
    return total_sum


if __name__ == "__main__":
    print(Part1(sys.stdin))
