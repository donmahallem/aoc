import codecs
import sys
import typing


def parseInput(inp: typing.TextIO) -> list[tuple[bool, int]]:
    return [(line[0] == "L", int(line[1:].strip())) for line in inp]


# parseInput Gen
def parseInputGen(inp: typing.TextIO) -> typing.Generator[int, None, None]:
    for line in inp:
        yield (-1 if line[0] == "L" else 1) * int(line[1:].strip())
