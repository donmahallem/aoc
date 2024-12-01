import codecs
import regex
import numpy as np

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]
search_terms = "XMAS"
zeilen = len(data)
spalten = len(data[0])


def checkDir(x, y, dirx, diry, term):
    if x + (dirx * len(term)) + 1 < 0 or x + (dirx * len(term)) > spalten:
        return False
    elif y + (diry * len(term)) + 1 < 0 or y + (diry * len(term)) > zeilen:
        return False
    for i in range(len(term)):
        if x == 3 and y == 9 and dirx == -1:
            print(data[y + diry * i][x + dirx * i], term[i])
        if data[y + diry * i][x + dirx * i] != term[i]:
            return False
    return True


count = 0
dirs = [[[-1, -1, 1, 1], [1, 1, -1, -1]], [[-1, 1, 1, -1], [1, -1, -1, 1]]]
for y in range(1, zeilen - 1):
    for x in range(1, spalten - 1):
        matches = 0
        if data[y][x] == "A":
            for axis in dirs:
                for dirx1, diry1, dirx2, diry2 in axis:
                    if (
                        data[y + diry1][x + dirx1] == "M"
                        and data[y + diry2][x + dirx2] == "S"
                    ):
                        matches += 1
                        # Only max one occurence per axis
                        break
        if matches == 2:
            count += 1
print(count)
