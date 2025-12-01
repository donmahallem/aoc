import codecs
import sys
import typing


def parseInput(inp: typing.TextIO) -> list[tuple[bool, int]]:
    return [(line[0] == 'L', int(line[1:].strip())) for line in inp]
