import sys
import typing
from .shared import _parse_line


def Part1(input: typing.TextIO) -> int:
    total = 0
    for line in input:
        line = line.strip()
        if not line:
            continue
        game_id, blocks = _parse_line(line)
        if all(b.red <= 12 and b.green <= 13 and b.blue <= 14 for b in blocks):
            total += game_id
    return total


if __name__ == "__main__":
    print(Part1(sys.stdin))
