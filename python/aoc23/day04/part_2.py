import sys
import typing
from .shared import parse_line, count_wins


def Part2(input: typing.TextIO) -> int:
    win_counts: list[int] = []
    for line in input:
        line = line.strip()
        if not line:
            continue
        _, winning, picked = parse_line(line)
        win_counts.append(count_wins(winning, picked))

    # Each card index starts with 1 copy; winning copies of subsequent cards
    copies = [1] * len(win_counts)
    for i, wins in enumerate(win_counts):
        for j in range(i + 1, min(i + 1 + wins, len(win_counts))):
            copies[j] += copies[i]

    return sum(copies)


if __name__ == "__main__":
    print(Part2(sys.stdin))
