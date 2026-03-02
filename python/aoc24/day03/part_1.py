import re
import sys
import typing


def Part1(input: typing.TextIO) -> int:
    data = input.read()
    pattern = re.compile(r"mul\((\d+),(\d+)\)")
    findings = pattern.findall(data)
    return sum(int(a) * int(b) for a, b in findings)


if __name__ == "__main__":
    Part1(sys.stdin)
