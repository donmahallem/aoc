import typing
import sys
import re


def Part2(input: typing.TextIO) -> int:
    data = input.readlines()
    data = "\n".join(data)
    enabled = True
    comp = re.compile(r"do(n\'t)?\(\)", flags=re.MULTILINE)
    comp2 = re.compile(r"mul\((\d+)\,(\d+)\)", flags=re.MULTILINE)
    total = 0
    for item in comp.split(data):
        if item == "n't":
            enabled = False
            continue
        elif item == None:
            enabled = True
        elif enabled:
            findings = comp2.findall(item)
            total += sum([int(a) * int(b) for a, b in findings])
    return total


if __name__ == "__main__":
    Part2(sys.stdin)
