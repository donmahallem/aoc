import typing
import numpy as np

def find_connected(field,checked_map,y, x, val, alread_connected):
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    for dir_y, dir_x in dirs:
        new_x, new_y = dir_x + x, dir_y + y
        if (
            new_x < 0
            or new_y < 0
            or new_x >= field.shape[1]
            or new_y >= field.shape[0]
        ):
            continue
        elif checked_map[new_y, new_x] == 1:
            continue
        elif (new_y, new_x) in alread_connected:
            continue
        if field[new_y, new_x] == val:
            alread_connected.add((new_y, new_x))
            checked_map[new_y, new_x] = 1
            find_connected(field,checked_map,new_y, new_x, val, alread_connected)


def count_edges(coords):
    edge_count = 0
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    for coord in coords:
        y, x = coord
        for dir_y, dir_x in dirs:
            new_x, new_y = dir_x + x, dir_y + y
            if (new_y, new_x) in coords:
                continue
            edge_count += 1
    return edge_count

def parseInput(input:typing.TextIO)->np.typing.NDArray:
    data = [[ord(num) for num in a.strip()] for a in input.readlines()]
    return np.array(data, dtype=np.uint)

def Part1(input: typing.TextIO) -> int:
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

    return sum([len(group) * count_edges(group) for _, group in groups])
