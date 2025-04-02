import typing
import numpy as np


def translateMovement(a):
    if a == "<":
        return (0, -1)
    elif a == ">":
        return (0, 1)
    elif a == "v":
        return (1, 0)
    elif a == "^":
        return (-1, 0)


CELL_BOX_LEFT = 4
CELL_BOX_RIGHT = 3
CELL_WALL = 2
CELL_BOX = 1
CELL_EMPTY = 0
