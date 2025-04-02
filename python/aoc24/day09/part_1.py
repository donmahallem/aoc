import codecs

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

data = [[int(num) for num in row] for row in data]

def handleRow(row):
    line = []
    block_num = 0
    for i in range(len(row)):
        if i % 2 == 0:
            for j in range(row[i]):
                line.append(block_num)
            block_num += 1
        else:
            for j in range(row[i]):
                line.append(-1)
    idx = 0
    while True:
        while True:
            if line[-1] == -1:
                line = line[:-1]
            else:
                break
        if line[idx] < 0:
            line[idx] = line[-1]
            line = line[:-1]
        idx += 1
        if idx >= len(line):
            break
    return line

for row in data:
    line_data = handleRow(row)
    print(sum([i * num for i, num in enumerate(line_data) if num >= 0]))
