import typing
import numpy as np
from .shared import shortestPath, calculatePathCost, parseField, CountSheats

def handle(input: typing.TextIO, cheatSavings: int) -> int:
    field, player_position, end_position = parseField(input)
    path_cost = calculatePathCost(field, player_position)
    normal_path_taken = shortestPath(field, path_cost, end_position)
    return CountSheats(normal_path_taken, path_cost, cheatSavings)

def Part1(input: typing.TextIO) -> int:
    return handle(input, 100)
