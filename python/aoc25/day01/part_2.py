import sys
import typing
from .parse_input import parseInput


def Part2(input: typing.TextIO) -> int:
    l = parseInput(input)

    current_position = 50
    zeros = 0
    for turn_left, distance in l:
        start = current_position
        if distance == 0:
            if start == 0:
                zeros += 1
            # position unchanged
            continue

        zero_hits = 0
        if turn_left:
            zero_hits = start % 100
        else:
            zero_hits = (100 - start) % 100
        if zero_hits == 0:
            zero_hits = 100
        if distance >= zero_hits:
            zeros += 1 + (distance - zero_hits) // 100

        if turn_left:
            current_position = (start - distance) % 100
        else:
            current_position = (start + distance) % 100
    return zeros


if __name__ == "__main__":
    Part2(sys.stdin)
