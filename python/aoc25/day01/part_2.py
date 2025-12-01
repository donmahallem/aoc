import sys
import typing
from .parse_input import parseInputGen


def Part2(inp: typing.TextIO) -> int:
    current_position = 50
    zeros = 0
    for distance in parseInputGen(inp):
        start = current_position
        if distance == 0:
            if start == 0:
                zeros += 1
            continue

        left = distance < 0
        steps = abs(distance)

        # first step t >= 1 that reaches 0
        first_hit = start % 100 if left else (100 - start) % 100
        if first_hit == 0:
            first_hit = 100
        if steps >= first_hit:
            zeros += 1 + (steps - first_hit) // 100

        current_position = (start - steps) % 100 if left else (start +
                                                               steps) % 100

    return zeros


if __name__ == "__main__":
    print(Part2(sys.stdin))
