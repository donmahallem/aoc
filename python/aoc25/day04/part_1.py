import sys
import typing
import numpy as np

surrounding = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0),
               (1, 1)]


def moveable(field: np.ndarray) -> typing.List[typing.Tuple[int, int]]:
    moves = []
    h, w = field.shape
    for y in range(h):
        for x in range(w):
            if not field[y, x]:
                continue
            count = 0
            for dy, dx in surrounding:
                ny = y + dy
                nx = x + dx
                if 0 <= ny < h and 0 <= nx < w:
                    if field[ny, nx]:
                        count += 1

            if count < 4:
                moves.append((y, x))
    return moves


def parseInput(input: typing.TextIO) -> np.ndarray:
    lines = [line.strip() for line in input]
    height = len(lines)
    width = len(lines[0])
    field = np.zeros((height, width), dtype=bool)
    for y in range(height):
        for x in range(width):
            if lines[y][x] == '@':
                field[y, x] = True
    return field


def Part1(input: typing.TextIO) -> int:
    field = parseInput(input)

    return len(moveable(field))


if __name__ == "__main__":
    print(Part1(sys.stdin))
