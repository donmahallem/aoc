import typing
from .shared import _parse_input


def Part2(
    input: typing.TextIO, width: int = 101, height: int = 103, maxSteps: int = 2000000
) -> int:
    pxs, pys, vxs, vys = _parse_input(input)

    num_robots = len(pxs)

    max_steps_to_check = width * height if (width * height) < maxSteps else maxSteps

    for i in range(1, max_steps_to_check + 1):
        positions = set()
        unique = True
        for j in range(num_robots):
            nx = (pxs[j] + vxs[j] * i) % width
            ny = (pys[j] + vys[j] * i) % height
            if (nx, ny) in positions:
                unique = False
                break
            positions.add((nx, ny))

        if unique:
            return i

    return -1
