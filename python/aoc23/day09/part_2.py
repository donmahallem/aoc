import sys
import typing
from .shared import _parse


def _predict_left(row: list[int]) -> int:
    """Return the previous value in the sequence."""
    row = list(row)
    n = len(row)
    for start in range(n - 1):
        all_zero = True
        for i in range(n - 1, start, -1):
            row[i] = row[i] - row[i - 1]
            if row[i] != 0:
                all_zero = False
        if all_zero:
            val = 0
            for up in range(start, -1, -1):
                val = row[up] - val
            return val
    return 0


def Part2(input: typing.TextIO) -> int:
    return sum(_predict_left(row) for row in _parse(input))


if __name__ == "__main__":
    print(Part2(sys.stdin))
