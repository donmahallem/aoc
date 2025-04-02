import typing
import sys
from .shared import checkRow

def Part2(input: typing.TextIO) -> int:
    data = input.readlines()

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
    return safe_count

if __name__ == "__main__":
    Part2(sys.stdin)
