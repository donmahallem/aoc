import typing
import sys


def Part1(input: typing.TextIO) -> int:
    data = "\n".join([a.strip() for a in input.readlines()])
    ordering_raw, pages_raw = data.split("\n\n")
    ordering = [
        tuple([int(a) for a in row.split("|")])
        for row in ordering_raw.split("\n")
    ]
    pages = [[int(a) for a in row.split(",")] for row in pages_raw.split("\n")]

    rules = dict()
    for a, b in ordering:
        if a not in rules:
            rules[a] = [b]
        else:
            rules[a].append(b)

    def validRow(row):
        valid = True
        for idx in range(len(row) - 1):
            if not (row[idx] in rules and row[idx + 1] in rules[row[idx]]):
                valid = False
                break
        return valid

    valid_pages = []
    for page in pages:
        if validRow(page):
            valid_pages.append(page)
    return sum([page[len(page) // 2] for page in valid_pages])


if __name__ == "__main__":
    Part1(sys.stdin)
