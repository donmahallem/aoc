import re
import sys
import typing

def Part1(input: typing.TextIO) -> int:
    data = input.readlines()
    data = "".join(data)
    comp = re.compile(r"mul\((\d+)\,(\d+)\)", flags=re.MULTILINE)
    findings = comp.findall(data)
    return sum([int(a) * int(b) for a, b in findings])

if __name__ == "__main__":
    Part1(sys.stdin)
