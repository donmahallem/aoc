import typing
from .part_1 import parseInput
import numpy as np


def isTree(robots):
    a = np.unique(robots[:, 0:2], axis=0).shape[0]
    return a == robots.shape[0]


def Part2(
    input: typing.TextIO, width: int = 101, height: int = 103, maxSteps: int = 2000000
) -> int:
    robots = parseInput(input)
    i = 0
    while i < maxSteps:
        robots[:, 0:2] = robots[:, 0:2] + robots[:, 2:4]
        robots[:, 0] = robots[:, 0] % width
        robots[:, 1] = robots[:, 1] % height
        if isTree(robots):
            return i
        i += 1
    return -1
