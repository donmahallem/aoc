import codecs
import re
import numpy as np

test_data = False
steps = 100
height = 7 if test_data else 103
width = 11 if test_data else 101
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.readlines()

parse_regex = re.compile(r"p=([+-]?\d+),([+-]?\d+).+?([+-]?\d+),([+-]?\d+)")
robots = np.zeros((len(data), 4), dtype=np.int64)
for i, machine in enumerate(data):
    reg_res = parse_regex.match(machine)
    if reg_res:
        robots[i, 0] = int(reg_res.groups()[0])
        robots[i, 1] = int(reg_res.groups()[1])
        robots[i, 2] = int(reg_res.groups()[2])
        robots[i, 3] = int(reg_res.groups()[3])

step = robots[:, 0:2] + robots[:, 2:4] * 100
step[:, 0] = step[:, 0] % width
step[:, 1] = step[:, 1] % height

center_x = width // 2
center_y = height // 2
q1 = (step[:, 0] < center_x) * (step[:, 1] < center_y)
q2 = (step[:, 0] > center_x) * (step[:, 1] < center_y)
q3 = (step[:, 0] > center_x) * (step[:, 1] > center_y)
q4 = (step[:, 0] < center_x) * (step[:, 1] > center_y)
print(sum(q1) * sum(q2) * sum(q3) * sum(q4))
