import codecs
import sys
import typing


# parseInput Gen
def _parse_input(inp: typing.TextIO) -> list[tuple[int, int]]:
    items: list[tuple[int, int]] = []
    for line in inp:
        for pair in line.strip().split(","):
            if pair == "":
                continue
            a, b = pair.split("-")
            items.append((int(a), int(b)))

    return items
