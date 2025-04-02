import typing
import numpy as np
from .part_1 import parseInput,find_connected

def count_straight_edges(field,coords):
    test = np.zeros((field.shape[0] + 2, field.shape[1] + 2), dtype=np.int8)
    for y, x in coords:
        test[y + 1, x + 1] = 1
    diff_0 = np.diff(test, axis=0)
    edges = 0
    for y in range(diff_0.shape[0]):
        last_val = 0
        for x in range(diff_0.shape[1]):
            if diff_0[y, x] != 0 and (last_val != diff_0[y, x]):
                edges += 1
                last_val = diff_0[y, x]
            elif diff_0[y, x] == 0 and last_val != 0:
                last_val = 0
    # horizontale Kanten
    diff_1 = np.diff(test, axis=1)
    for x in range(diff_1.shape[1]):
        last_val = 0
        for y in range(diff_1.shape[0]):
            if diff_1[y, x] != 0 and (last_val != diff_1[y, x]):
                edges += 1
                last_val = diff_1[y, x]
            elif diff_1[y, x] == 0 and last_val != 0:
                last_val = 0
    return edges


def Part2(input: typing.TextIO) -> int:
    field=parseInput(input)
    checked_map = np.zeros((field.shape), dtype=np.uint8)
    groups = list()
    for y in range(field.shape[0]):
        for x in range(field.shape[1]):
            if checked_map[y, x] == 1:
                continue
            group = set({(y, x)})
            find_connected(field,checked_map,y, x, field[y, x], group)
            key = chr(field[y, x])
            groups.append((key, group))
    count = 0
    for key, group in groups:
        count += count_straight_edges(field,group) * len(group)
    return count
