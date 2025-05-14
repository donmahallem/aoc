import typing
import sys


def Part2(input: typing.TextIO) -> int:
    raw_lines = "\n".join([a.strip() for a in input.readlines()])
    ordering_raw, pages_raw = raw_lines.split("\n\n")
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

    def validRow(row: list[int]):
        valid = True
        for idx in range(len(row) - 1):
            if not (row[idx] in rules and row[idx + 1] in rules[row[idx]]):
                valid = False
                break
        return valid

    def sort(items):
        invalid = True
        while invalid:
            invalid = False
            for n in range(len(items) - 1, 0, -1):
                for i in range(0, n):
                    if items[n] in rules and items[i] in rules[items[n]]:
                        items[n], items[i] = items[i], items[n]
                        invalid = True

    invalid_pages: list[list[int]] = []
    for page in pages:
        if not validRow(page):
            invalid_pages.append(page)

    for i in range(len(invalid_pages)):
        sort(invalid_pages[i])
    return sum([page[len(page) // 2] for page in invalid_pages])


if __name__ == "__main__":
    Part2(sys.stdin)
