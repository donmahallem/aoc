import typing
import sys


def Part1(input: typing.TextIO) -> int:
    data = [a.strip() for a in input.readlines()]
    data = "\n".join(data)
    ordering, pages = data.split("\n\n")
    ordering = [tuple([int(a) for a in row.split("|")]) for row in ordering.split("\n")]
    pages = [[int(a) for a in row.split(",")] for row in pages.split("\n")]

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
