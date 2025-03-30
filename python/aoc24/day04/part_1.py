import sys
import typing


def Part1(input: typing.TextIO) -> int:
    data = [a.strip() for a in input.readlines()]
    search_terms = "XMAS"
    zeilen = len(data)
    spalten = len(data[0])
    count = 0

    def checkDir(x, y, dirx, diry, term):
        if x + (dirx * len(term)) + 1 < 0 or x + (dirx * len(term)) > spalten:
            return False
        elif y + (diry * len(term)) + 1 < 0 or y + (diry * len(term)) > zeilen:
            return False
        for i in range(len(term)):
            if data[y + diry * i][x + dirx * i] != term[i]:
                return False
        return True

    count = 0
    dirs = [[1, 0], [1, -1], [0, -1], [-1, -1], [-1, 0], [-1, 1], [0, 1], [1, 1]]
    for y in range(0, zeilen):
        for x in range(spalten):
            for xdir, ydir in dirs:
                if checkDir(x, y, xdir, ydir, search_terms):
                    count += 1
    return count


if __name__ == "__main__":
    Part1(sys.stdin)
