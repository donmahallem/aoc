import sys
import typing
from .parse_input import __parseInput


def Part1(input: typing.TextIO) -> int:
    start, splitter, width, height = __parseInput(input)
    splitEvents: int = 0
    beams: list[int] = [start[0] + start[1] * width]
    visited: dict[int, bool] = dict()
    while len(beams) > 0:
        beam = beams.pop()
        if beam in visited:
            continue
        visited[beam] = True
        x = beam % width
        y = beam // width
        if y >= height:
            continue
        if beam in splitter:
            splitEvents += 1
            if x > 0:
                beams.append(beam - 1 + width)
            if x < width - 1:
                beams.append(beam + 1 + width)
        else:
            beams.append(beam + width)
    return splitEvents


if __name__ == "__main__":
    print(Part1(sys.stdin))
