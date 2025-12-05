import sys
import typing
from aoc25.day04.part_1 import moveable, parseInput, surrounding


def is_moveable(field, y: int, x: int) -> bool:
    if not field[y, x]:
        return False
    h, w = field.shape
    count = 0
    for dy, dx in surrounding:
        ny = y + dy
        nx = x + dx
        if 0 <= ny < h and 0 <= nx < w and field[ny, nx]:
            count += 1
            if count >= 4:
                return False
    return True


def Part2(input: typing.TextIO) -> int:
    field = parseInput(input)

    initial = moveable(field)
    stack: typing.Set[typing.Tuple[int, int]] = set(initial)

    removed = 0
    h, w = field.shape

    while stack:
        y, x = stack.pop()
        # skip if already gone
        if not field[y, x]:
            continue

        field[y, x] = False
        removed += 1

        for dy, dx in surrounding:
            ny = y + dy
            nx = x + dx
            if 0 <= ny < h and 0 <= nx < w and field[ny, nx]:
                if is_moveable(field, ny, nx):
                    stack.add((ny, nx))

    return removed


if __name__ == "__main__":
    print(Part2(sys.stdin))
