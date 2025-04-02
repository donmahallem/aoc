import codecs
import numpy as np

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

data = [[int(num) for num in row] for row in data][0]
largest_num = max(data)

def expandToGroups(data):
    line = []
    block_num = 0
    for end_idx in range(len(data)):
        if end_idx % 2 == 0:
            line.append((block_num, data[end_idx]))
            block_num += 1
        elif data[end_idx] > 0:
            line.append((-1, data[end_idx]))
    return (line, block_num - 1)

def handleRow(line, current_idx):
    for end_idx in range(len(line) - 1, 0, -1):
        if line[end_idx][0] != current_idx:
            continue
        for start_idx in range(0, end_idx):
            if line[start_idx][0] != -1:
                continue
            # Matching group size
            if line[start_idx][1] == line[end_idx][1]:
                line[start_idx] = line[end_idx]
                line[end_idx] = (-1, line[end_idx][1])
                break
            # Spaces larger than group
            elif line[start_idx][1] > line[end_idx][1]:
                line = (
                    line[0:start_idx]
                    + [line[end_idx], (-1, line[start_idx][1] - line[end_idx][1])]
                    + line[start_idx + 1 : end_idx]
                    + [(-1, line[end_idx][1])]
                    + line[end_idx + 1 :]
                )
                runner = 0
                while True:
                    if runner >= len(line) - 1:
                        break
                    if line[runner][0] == -1 and line[runner + 1][0] == -1:
                        line[runner] = (-1, line[runner + 1][1] + line[runner][1])
                        line.pop(runner + 1)
                    else:
                        runner += 1
                break
        break
    return line

line_data, max_num = expandToGroups(data)

for i in range(max_num, 0, -1):
    line_data = handleRow(line_data, i)
line_data = [a for a, b in line_data for x in range(0, b)]
print(sum([i * num for i, num in enumerate(line_data) if num >= 0]))
