import sys
import typing
import math
from .shared import _parse_input


def _get_cycle_size(
    instructions: list[int], nodes: dict[str, list[str]], start: str
) -> int:
    """Return cycle length from start until we hit a **Z node aligned with instruction modulo."""
    n = len(instructions)
    current = start
    i = 0
    while True:
        direction = instructions[i % n]
        current = nodes[current][direction]
        i += 1
        if current.endswith("Z") and i % n == 0:
            return i


def Part2(input: typing.TextIO) -> int:
    instructions, nodes = _parse_input(input)
    starts = [node for node in nodes if node.endswith("A")]
    result = 1
    for start in starts:
        cycle = _get_cycle_size(instructions, nodes, start)
        result = math.lcm(result, cycle)
    return result


if __name__ == "__main__":
    print(Part2(sys.stdin))
