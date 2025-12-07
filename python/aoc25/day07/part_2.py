import sys
import typing
from .parse_input import __parseInput


def Part2(input: typing.TextIO) -> int:
    (x, y), splitter, width, height = __parseInput(input)
    cache: dict[int, int] = dict()

    def dfs(beamX, beamY):
        stack: list[tuple[int, int]] = []
        stack.append((beamX, beamY))
        path = []
        while stack:
            bx, by = stack.pop()
            key = (bx, by)
            if key in cache:
                path.append(cache[key])
                continue
            if by == height:
                cache[key] = 1
                path.append(1)
                continue
            beamPos = bx + by * width
            total = 0
            if beamPos in splitter:
                if bx > 0:
                    total += dfs(bx - 1, by + 1)
                if bx < width - 1:
                    total += dfs(bx + 1, by + 1)
            else:
                total += dfs(bx, by + 1)
            cache[key] = total
            path.append(total)
        return path[-1] if path else 0

    return dfs(x, y)


if __name__ == "__main__":
    print(Part2(sys.stdin))
