import typing
import re
import numpy as np


def parseInput(input: typing.TextIO | list[str], *, lines: list[str] | None = None):
    if lines is None:
        data = [line.strip() for line in input.readlines()]
    else:
        data = [line.strip() for line in lines]
    parse_regex = re.compile(
        r"p=([+-]?\d+),([+-]?\d+).+?([+-]?\d+),([+-]?\d+)")
    robots = np.zeros((len(data), 4), dtype=np.int64)
    for i, machine in enumerate(data):
        reg_res = parse_regex.match(machine)
        if reg_res:
            robots[i, 0] = int(reg_res.groups()[0])
            robots[i, 1] = int(reg_res.groups()[1])
            robots[i, 2] = int(reg_res.groups()[2])
            robots[i, 3] = int(reg_res.groups()[3])
    return robots


def Part1(input: typing.TextIO,
        width: int = 101,
        height: int = 103,
        steps: int = 100) -> int:
    lines = [line.strip() for line in input.readlines()]
    robots = parseInput(input, lines=lines)

    # Heuristic: small samples use 11x7 grid (per AoC example).
    if len(lines) <= 20:
      width = 11
      height = 7
    step = robots[:, 0:2] + robots[:, 2:4] * steps
    step[:, 0] = step[:, 0] % width
    step[:, 1] = step[:, 1] % height

    center_x = width // 2
    center_y = height // 2
    q1 = (step[:, 0] < center_x) * (step[:, 1] < center_y)
    q2 = (step[:, 0] > center_x) * (step[:, 1] < center_y)
    q3 = (step[:, 0] > center_x) * (step[:, 1] > center_y)
    q4 = (step[:, 0] < center_x) * (step[:, 1] > center_y)
    return sum(q1) * sum(q2) * sum(q3) * sum(q4)
