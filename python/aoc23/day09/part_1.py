import sys
import typing
from .shared import _parse


def _predict_right(row: list[int]) -> int:
    """Return the next value in the sequence."""
    row = list(row)
    n = len(row)
    cum_sum = 0
    for end in range(n, 0, -1):
        all_zero = True
        for i in range(end - 1):
            row[i] = row[i + 1] - row[i]
            if row[i] != 0:
                all_zero = False
        if all_zero:
            for idx in range(end - 1, n):
                cum_sum += row[idx]
            return cum_sum
    return 0


def Part1(input: typing.TextIO) -> int:
    return sum(_predict_right(row) for row in _parse(input))


if __name__ == "__main__":
    print(Part1(sys.stdin))
