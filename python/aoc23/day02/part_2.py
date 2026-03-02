import sys
import typing
from .shared import parse_line


def Part2(input: typing.TextIO) -> int:
    total = 0
    for line in input:
        line = line.strip()
        if not line:
            continue
        _, blocks = parse_line(line)
        min_red = max(b.red for b in blocks)
        min_green = max(b.green for b in blocks)
        min_blue = max(b.blue for b in blocks)
        total += min_red * min_green * min_blue
    return total


if __name__ == "__main__":
    print(Part2(sys.stdin))
