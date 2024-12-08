import codecs
import numpy as np

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

height = len(data)
width = len(data[0])
base_field = np.zeros((height, width), dtype=np.uint8)
for y in range(len(data)):
    for x in range(len(data[0])):
        if data[y][x] == ".":
            continue
        else:
            base_field[y, x] = ord(data[y][x])

uni = np.delete(np.unique(base_field), 0)

occurences = {key: list(zip(*np.where(base_field == key))) for key in uni}


def inside(y, x):
    return x >= 0 and y >= 0 and x < width and y < height


antinodes = set()
for key in occurences.keys():
    nodes = occurences[key]
    antinodes.update(nodes)
    for i in range(0, len(nodes) - 1):
        for j in range(i + 1, len(nodes)):
            diff = (nodes[i][0] - nodes[j][0], nodes[i][1] - nodes[j][1])
            for u in range(1, 2000):
                if inside(nodes[i][0] + (u * diff[0]), nodes[i][1] + (u * diff[1])):
                    antinodes.add(
                        (nodes[i][0] + (u * diff[0]), nodes[i][1] + (u * diff[1]))
                    )
                else:
                    break
            for u in range(1, 2000):
                if inside(nodes[j][0] - (u * diff[0]), nodes[j][1] - (u * diff[1])):
                    antinodes.add(
                        (nodes[j][0] - (u * diff[0]), nodes[j][1] - (u * diff[1]))
                    )
                else:
                    break
print(len(antinodes))
