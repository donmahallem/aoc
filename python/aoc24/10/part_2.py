import codecs
import numpy as np

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]
data = [[int(item) for item in line.strip()] for line in data]

data_np = np.array(data, dtype=np.uint8)
trailheads = list(zip(*np.where(data_np == 0)))


def walk(y, x, looking_for, ends):
    sum = 0
    dirs = [(0, 1), (1, 0), (-1, 0), (0, -1)]
    for dir_y, dir_x in dirs:
        check_x, check_y = dir_x + x, dir_y + y
        if (
            check_x < 0
            or check_y < 0
            or check_y >= data_np.shape[0]
            or check_x >= data_np.shape[1]
        ):
            continue
        if looking_for == 9 and data_np[check_y, check_x] == 9:
            sum += 1
            ends.add((check_y, check_x))
        elif data_np[check_y, check_x] == looking_for:
            sum += walk(check_y, check_x, looking_for + 1, ends)
    return sum


summe = 0
for trailhead in trailheads:
    k = set()
    summe += walk(trailhead[0], trailhead[1], 1, k)
    # summe+=len(k)
print(summe)
