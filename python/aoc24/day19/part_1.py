import typing
from functools import cache
from .shared import parseInput

def Part1(input: typing.TextIO) -> int:
    available_patterns, needed_patterns = parseInput(input)

    @cache
    def checkPattern2(pattern):
        for test_pattern in available_patterns:
            if test_pattern == pattern:
                return True
            elif pattern[0 : len(test_pattern)] == test_pattern:
                if checkPattern2(pattern[len(test_pattern) :]):
                    return True
        return False

    results = 0
    for p in needed_patterns:
        if checkPattern2(p):
            results += 1

    return results
