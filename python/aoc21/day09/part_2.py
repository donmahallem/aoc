import sys
import typing

from .shared import parse_input


def Part2(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    count = 0
    centers = []
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
                centers.append((y, x))

    # Calculate the sizes of the basins
    basin_sizes = []
    for cy, cx in centers:
        visited = set()
        stack = [(cy, cx)]
        while stack:
            y, x = stack.pop()
            if (y, x) in visited:
                continue
            visited.add((y, x))
            for dy, dx in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                ny, nx = y + dy, x + dx
                if (
                    0 <= ny < len(data)
                    and 0 <= nx < len(data[ny])
                    and data[ny][nx] != 9
                ):
                    stack.append((ny, nx))
        basin_sizes.append(len(visited))
    basin_sizes.sort(reverse=True)
    if len(basin_sizes) >= 3:
        count = basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
    return count


if __name__ == "__main__":
    print(Part2(sys.stdin))
