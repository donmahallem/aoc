import sys
import typing
import re


def __parseInput(input: typing.TextIO) -> tuple[list[list[int]], list[str]]:
    pattern = re.compile(
        r'^(?P<numbers>\s*\d+(?:\s+\d+)+\s*)$'  # lines with numbers
        r'|^(?P<operators>\s*[+*](?:\s+[+*])*\s*)$',  # lines with +/*
        re.MULTILINE)

    nums: list[list[int]] = list()
    ops: list[str] = list()
    for line in input:
        line = line.strip()
        if line == "":
            continue
        match = pattern.match(line)
        if match:
            if match.group('numbers'):
                numsLine = match.group('numbers').strip().split()
                nums.append([int(n) for n in numsLine])
            elif match.group('operators'):
                opsLine = match.group('operators').strip().split()
                ops.extend(opsLine)
                break
    return nums, ops


def Part1(input: typing.TextIO) -> int:
    numRows, operators = __parseInput(input)
    output = numRows[0]
    for i in range(1, len(numRows)):
        output = [
            a + c if b == '+' else a * c
            for (a, b, c) in zip(output, operators, numRows[i])
        ]
    return sum(output)


if __name__ == "__main__":
    print(Part1(sys.stdin))
