import typing
import numpy as np
from .shared import shortestPath, calculatePathCost, parseField, CountSheats

def handle(input: typing.TextIO, cheatSavings: int) -> int:
    field, player_position, end_position = parseField(input)
    path_cost = calculatePathCost(field, player_position)
    normal_path_taken = shortestPath(field, path_cost, end_position)

    summe = 0
    for step_a in range(0, len(normal_path_taken) - 1):
        for step_b in range(len(normal_path_taken) - 1, step_a, -1):
            p1_y, p1_x = normal_path_taken[step_a]
            p2_y, p2_x = normal_path_taken[step_b]

            dst = abs(p1_x - p2_x) + abs(p1_y - p2_y)
            if dst < 2 or dst > 20:
                continue
            saved = step_b - step_a - dst
            if path_cost[p1_y, p1_y] - path_cost[p1_y, p1_y] == saved:
                continue
            elif saved < cheatSavings:
                continue
            summe += 1
    return summe

def Part2(input: typing.TextIO) -> int:
    return handle(input, 100)
