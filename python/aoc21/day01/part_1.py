import re
import sys
import typing
import math


def Part1(input_data: typing.TextIO) -> int:
    increased: int = 0
    lastValue: int | None = None
    for line in input_data.read().splitlines():
        if not line:
            continue
        value = int(line)
        if lastValue is not None and value > lastValue:
            increased += 1
        lastValue = value
    return increased


if __name__ == "__main__":
    print(Part1(sys.stdin))
