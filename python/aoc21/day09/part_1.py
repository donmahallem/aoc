import sys
import typing

from .shared import parse_input


def Part1(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    count = 0
    for y in range(len(data)):
        for x in range(len(data[y])):
            value = data[y][x]
            # Check if the current value is a low point
            is_low_point = True
            for dy, dx in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                ny, nx = y + dy, x + dx
                if (
                    0 <= ny < len(data)
                    and 0 <= nx < len(data[ny])
                    and data[ny][nx] <= value
                ):
                    is_low_point = False
                    break
            if is_low_point:
                count += value + 1

    return count


if __name__ == "__main__":
    print(Part1(sys.stdin))
