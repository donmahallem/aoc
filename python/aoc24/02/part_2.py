import codecs

with codecs.open("data_1.txt", encoding="utf8") as f:
    data = f.readlines()


def checkRow(row):
    inc = None
    for i in range(0, len(row) - 1):
        if row[i] == row[i + 1]:
            return False
        if abs(row[i] - row[i + 1]) > 3:
            return False
        if not inc is None:
            if row[i] < row[i + 1] and not inc:
                return False
            if row[i] > row[i + 1] and inc:
                return False
        else:
            inc = row[i] < row[i + 1]

    return True


safe_count = 0
for dataline in data:
    row_data = [int(d) for d in dataline.split()]
    if not checkRow(row_data):
        for i in range(0, len(row_data)):
            if checkRow(row_data[0:i] + row_data[i + 1 :]):
                safe_count += 1
                break
    else:
        safe_count += 1
print(safe_count)
