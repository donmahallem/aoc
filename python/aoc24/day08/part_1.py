import typing
from .shared import parseInput, inside

def Part1(input: typing.TextIO) -> int:
    occurences, width, height = parseInput(input)
    antinodes = set()
    for key in occurences.keys():
        nodes = occurences[key]
        for i in range(0, len(nodes) - 1):
            for j in range(i + 1, len(nodes)):
                diff = (nodes[i][0] - nodes[j][0], nodes[i][1] - nodes[j][1])
                if inside(nodes[i][0] + diff[0], nodes[i][1] + diff[1], width, height):
                    antinodes.add((nodes[i][0] + diff[0], nodes[i][1] + diff[1]))
                if inside(nodes[j][0] - diff[0], nodes[j][1] - diff[1], width, height):
                    antinodes.add((nodes[j][0] - diff[0], nodes[j][1] - diff[1]))
    return len(antinodes)
