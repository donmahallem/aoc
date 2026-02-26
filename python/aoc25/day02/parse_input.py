import codecs
import sys
import typing


# parseInput Gen
def parseInputGen(
        inp: typing.TextIO) -> typing.Generator[tuple[int, int], None, None]:
    for line in inp:
        for pair in line.strip().split(','):
            if pair == '':
                continue
            a, b = pair.split('-')
            yield (int(a), int(b))
