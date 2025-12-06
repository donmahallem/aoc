import sys
import typing
import re


def parseInput(
        input: typing.TextIO) -> tuple[list[tuple[int, int]], list[str]]:
    numsParser = re.compile(r'[\d]+')
    operatorParser = re.compile(r'[+*]')
    nums = list()
    ops = list()
    for line in input:
        line = line.strip()
        if line == "":
            continue
        numsLine = numsParser.findall(line)
        if len(numsLine) == 0:
            opsLine = operatorParser.findall(line)
            if len(opsLine) > 0:
                ops.extend(opsLine)
                break
        else:
            nums.append([int(n) for n in numsLine])
    return nums, ops


def Part1(input: typing.TextIO) -> int:
    numRows, operators = parseInput(input)
    output = numRows[0]
    for i in range(1, len(numRows)):
        output = [
            a + c if b == '+' else a * c
            for (a, b, c) in zip(output, operators, numRows[i])
        ]
    return sum(output)


if __name__ == "__main__":
    print(Part1(sys.stdin))
