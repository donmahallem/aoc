import codecs
import math

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]

data = [row.split(":") for row in data]
data = [
    (int(result[0].strip()), [int(a) for a in result[1].strip().split(" ")])
    for result in data
]


def isValidRow(row):
    def nextOp(value, target, remaining_terms):
        if len(remaining_terms) == 0:
            return target == value
        elif value > target:
            return False
        return (
            nextOp(value + remaining_terms[0], target, remaining_terms[1:])
            or nextOp(value * remaining_terms[0], target, remaining_terms[1:])
            or nextOp(
                value * (10 ** (int(math.log10(remaining_terms[0])) + 1))
                + remaining_terms[0],
                target,
                remaining_terms[1:],
            )
        )

    result, terms = row
    return nextOp(terms[0], result, terms[1:])


print(sum([row[0] for row in data if isValidRow(row)]))
