import sys
import typing
from .part_1 import _count_options


def _parse_input2(input: typing.TextIO) -> tuple[int, int]:
    time: int = 0
    line: str = input.readline().strip()
    offset = line.index(" ") + 1
    for part in line[offset:]:
        if part.isdigit():
            time = time * 10 + int(part)
    distance: int = 0
    line = input.readline().strip()
    offset = line.index(" ") + 1
    for part in line[offset:]:
        if part.isdigit():
            distance = distance * 10 + int(part)
    return time, distance


def Part2(input: typing.TextIO) -> int:
    time, distance = _parse_input2(input)
    return _count_options(time, distance)


if __name__ == "__main__":
    print(Part2(sys.stdin))
