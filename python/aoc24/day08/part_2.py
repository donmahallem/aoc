import typing
from .shared import parseInput, inside, Position


def Part2(input: typing.TextIO) -> int:
    occurences, width, height = parseInput(input)

    antinodes: set[Position] = set()
    for key in occurences.keys():
        nodes = occurences[key]
        antinodes.update(nodes)
        for i in range(0, len(nodes) - 1):
            for j in range(i + 1, len(nodes)):
                diff = (nodes[i][0] - nodes[j][0], nodes[i][1] - nodes[j][1])
                for u in range(1, max(width, height)):
                    if inside(
                            nodes[i][0] + (u * diff[0]),
                            nodes[i][1] + (u * diff[1]),
                            width,
                            height,
                    ):
                        antinodes.add((nodes[i][0] + (u * diff[0]),
                                       nodes[i][1] + (u * diff[1])))
                    else:
                        break
                for u in range(1, max(width, height)):
                    if inside(
                            nodes[j][0] - (u * diff[0]),
                            nodes[j][1] - (u * diff[1]),
                            width,
                            height,
                    ):
                        antinodes.add((nodes[j][0] - (u * diff[0]),
                                       nodes[j][1] - (u * diff[1])))
                    else:
                        break
    return len(antinodes)
