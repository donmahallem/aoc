import codecs
import math

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

height = len(data)
width = len(data[0])

occurences = dict()
for y in range(len(data)):
    for x in range(len(data[0])):
        if data[y][x] == ".":
            continue
        if data[y][x] in occurences:
            occurences[data[y][x]].append((y, x))
        else:
            occurences[data[y][x]] = list([(y, x)])


def inside(y, x):
    return x >= 0 and y >= 0 and x < width and y < height


antinodes = set()
for key in occurences.keys():
    nodes = occurences[key]
    antinodes.update(nodes)
    for i in range(0, len(nodes) - 1):
        for j in range(i + 1, len(nodes)):
            diff = (nodes[i][0] - nodes[j][0], nodes[i][1] - nodes[j][1])
            for u in range(1, max(width, height)):
                if inside(nodes[i][0] + (u * diff[0]), nodes[i][1] + (u * diff[1])):
                    antinodes.add(
                        (nodes[i][0] + (u * diff[0]), nodes[i][1] + (u * diff[1]))
                    )
                else:
                    break
            for u in range(1, max(width, height)):
                if inside(nodes[j][0] - (u * diff[0]), nodes[j][1] - (u * diff[1])):
                    antinodes.add(
                        (nodes[j][0] - (u * diff[0]), nodes[j][1] - (u * diff[1]))
                    )
                else:
                    break
print(len(antinodes))
