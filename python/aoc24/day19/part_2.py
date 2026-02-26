import typing
from functools import cache
from .shared import parseInput


def Part2(input: typing.TextIO) -> int:
    available_patterns, needed_patterns = parseInput(input)

    @cache
    def checkPattern(pattern):
        valid_patterns = 0
        for test_pattern in available_patterns:
            if test_pattern == pattern:
                valid_patterns += 1
            elif pattern[0 : len(test_pattern)] == test_pattern:
                valid_patterns += checkPattern(pattern[len(test_pattern) :])
        return valid_patterns

    summe = 0
    for p in needed_patterns:
        summe += checkPattern(p)
    return summe
