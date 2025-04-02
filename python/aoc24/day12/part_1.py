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


groups = list()
for y in range(data_np.shape[0]):
    for x in range(data_np.shape[1]):
        if checked_map[y, x] == 1:
            continue
        group = set({(y, x)})
        find_connected(y, x, data_np[y, x], group)
        key = chr(data_np[y, x])
        groups.append((key, group))

print(sum([len(group) * count_edges(group) for key, group in groups]))
