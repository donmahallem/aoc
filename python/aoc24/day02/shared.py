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
