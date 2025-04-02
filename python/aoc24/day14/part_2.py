import codecs
import re
import numpy as np

test_data = False
height = 7 if test_data else 103
width = 11 if test_data else 101
steps = 100
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

steps = 100

def isTree(robots):
    a = np.unique(robots[:, 0:2], axis=0).shape[0]
    return a == robots.shape[0]

i = 0
while True:
    robots[:, 0:2] = robots[:, 0:2] + robots[:, 2:4]
    robots[:, 0] = robots[:, 0] % width
    robots[:, 1] = robots[:, 1] % height
    i += 1
    if isTree(robots):
        break
print(i)
