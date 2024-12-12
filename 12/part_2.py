import codecs
import numpy as np


with codecs.open("data.txt", encoding="utf8") as f:
    data = [[ord(num) for num in a.strip()] for a in f.readlines()]
data_np = np.array(data, dtype=np.uint)
checked_map = np.zeros((data_np.shape), dtype=np.uint8)


def find_connected(y, x, val, alread_connected):
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    for dir_y, dir_x in dirs:
        new_x, new_y = dir_x + x, dir_y + y
        if (
            new_x < 0
            or new_y < 0
            or new_x >= data_np.shape[1]
            or new_y >= data_np.shape[0]
        ):
            continue
        elif checked_map[new_y, new_x] == 1:
            continue
        elif (new_y, new_x) in alread_connected:
            continue
        if data_np[new_y, new_x] == val:
            alread_connected.add((new_y, new_x))
            checked_map[new_y, new_x] = 1
            find_connected(new_y, new_x, val, alread_connected)


def count_straight_edges(coords):
    test = np.zeros((data_np.shape[0] + 2, data_np.shape[1] + 2), dtype=np.int8)
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


groups = list()
for y in range(data_np.shape[0]):
    for x in range(data_np.shape[1]):
        if checked_map[y, x] == 1:
            continue
        group = set({(y, x)})
        find_connected(y, x, data_np[y, x], group)
        key = chr(data_np[y, x])
        groups.append((key, group))
summe = 0
for key, group in groups:
    summe += count_straight_edges(group) * len(group)
print(summe)
