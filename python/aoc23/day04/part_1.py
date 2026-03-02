import sys
import typing
from .shared import parse_line, count_wins


def Part1(input: typing.TextIO) -> int:
    total = 0
    for line in input:
        line = line.strip()
        if not line:
            continue
        _, winning, picked = parse_line(line)
        wins = count_wins(winning, picked)
        if wins > 0:
            total += 1 << (wins - 1)
    return total


if __name__ == "__main__":
    print(Part1(sys.stdin))
