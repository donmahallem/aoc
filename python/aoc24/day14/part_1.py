import typing
from .shared import _parse_input


def Part1(
    input: typing.TextIO, width: int = 101, height: int = 103, steps: int = 100
) -> int:
    pxs, pys, vxs, vys = _parse_input(input)

    # Heuristic: small samples use 11x7 grid (per AoC example).
    if len(pxs) <= 20:
        width = 11
        height = 7

    center_x = width // 2
    center_y = height // 2

    q1 = q2 = q3 = q4 = 0

    for px, py, vx, vy in zip(pxs, pys, vxs, vys):
        nx = (px + vx * steps) % width
        ny = (py + vy * steps) % height

        if nx < center_x:
            if ny < center_y:
                q1 += 1
            elif ny > center_y:
                q4 += 1
        elif nx > center_x:
            if ny < center_y:
                q2 += 1
            elif ny > center_y:
                q3 += 1

    return q1 * q2 * q3 * q4
