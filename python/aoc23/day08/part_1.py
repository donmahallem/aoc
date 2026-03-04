import sys
import typing
from .shared import _parse_input


def Part1(input: typing.TextIO) -> int:
    instructions, nodes = _parse_input(input)
    n = len(instructions)
    current = "AAA"
    steps = 0
    while current != "ZZZ":
        direction = instructions[steps % n]
        current = nodes[current][direction]
        steps += 1
    return steps


if __name__ == "__main__":
    print(Part1(sys.stdin))
