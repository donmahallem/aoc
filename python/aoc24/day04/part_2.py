import typing
import sys


def Part2(input: typing.TextIO) -> int:
    data = [a.strip() for a in input.readlines()]
    zeilen = len(data)
    spalten = len(data[0])
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
    return count


if __name__ == "__main__":
    Part2(sys.stdin)
