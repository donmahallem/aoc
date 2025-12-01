import sys
import typing
from .parse_input import parseInputGen


def Part1(input: typing.TextIO) -> int:
    

    current_position = 50
    zeros = 0
    for  distance in parseInputGen(input):
        current_position = (current_position + distance) % 100
        if current_position == 0:
            zeros += 1
    return zeros


if __name__ == "__main__":
    Part1(sys.stdin)
