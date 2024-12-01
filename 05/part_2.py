import codecs
import regex
import numpy as np
from functools import cmp_to_key

with codecs.open("data.txt", encoding="utf8") as f:
    data = [a.strip() for a in f.readlines()]
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


def sort(items):
    invalid = True
    while invalid:
        invalid = False
        for n in range(len(items) - 1, 0, -1):
            for i in range(0, n):
                if items[n] in rules and items[i] in rules[items[n]]:
                    items[n], items[i] = items[i], items[n]
                    invalid = True


invalid_pages = []
for page in pages:
    if not validRow(page):
        invalid_pages.append(page)

for i in range(len(invalid_pages)):
    sort(invalid_pages[i])
print(sum([page[len(page) // 2] for page in invalid_pages]))
