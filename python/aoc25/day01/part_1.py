import codecs
import sys
import typing
from .parse_input import parseInput


def Part1(input: typing.TextIO) -> int:
    l = parseInput(input)

    current_position = 50
    zeros = 0
    for turn_left, distance in l:
        if turn_left:
            current_position = (current_position - distance) % 100
        else:
            current_position = (current_position + distance) % 100
        if current_position == 0:
            zeros += 1
    return zeros


if __name__ == "__main__":
    Part1(sys.stdin)
