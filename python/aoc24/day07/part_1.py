import typing
from .shared import parseRows

def Part1(input: typing.TextIO) -> int:
    rows = parseRows(input)

    def isValidRow(row):
        def nextOp(value, target, remaining_terms):
            if len(remaining_terms) == 0:
                return target == value
            elif value > target:
                return False
            return nextOp(
                value + remaining_terms[0], target, remaining_terms[1:]
            ) or nextOp(value * remaining_terms[0], target, remaining_terms[1:])

        result, terms = row
        return nextOp(terms[0], result, terms[1:])

    return sum([row[0] for row in rows if isValidRow(row)])
