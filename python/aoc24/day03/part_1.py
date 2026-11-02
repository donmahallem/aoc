import re
import sys
import typing


def Part1(input: typing.TextIO) -> int:
    data = input.read()
    pattern = re.compile(r"mul\((\d{1,3}),(\d{1,3})\)")
    return sum([int(a) * int(b) for a, b in pattern.findall(data)])


if __name__ == "__main__":
    Part1(sys.stdin)
