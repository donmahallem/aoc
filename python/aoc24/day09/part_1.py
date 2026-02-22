import typing


def parseInput(input: typing.TextIO) -> list[int]:
    # single line of digits
    return [int(ch) for ch in input.readline().strip()]


def handleRow(row: list[int]) -> list[int]:
    line: list[int] = []
    block_num = 0
    for i, count in enumerate(row):
        if i % 2 == 0:
            if count:
                line.extend([block_num])
                if count > 1:
                    line.extend([block_num] * (count - 1))
            block_num += 1
        else:
            if count:
                line.extend([-1])
                if count > 1:
                    line.extend([-1] * (count - 1))
    end = len(line) - 1
    while end >= 0 and line[end] == -1:
        end -= 1
    i = 0
    while i <= end:
        if line[i] == -1:
            line[i] = line[end]
            end -= 1
            while end >= 0 and line[end] == -1:
                end -= 1
        i += 1
    # return the valid prefix
    return line[:end + 1]


def Part1(input: typing.TextIO) -> int:
    data = parseInput(input)
    line_data = handleRow(data)
    total = 0
    for idx, num in enumerate(line_data):
        if num >= 0:
            total += idx * num
    return total
