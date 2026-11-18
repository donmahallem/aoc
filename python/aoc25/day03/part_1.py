import sys
import typing


def reduce(s: str) -> int:
    tens = max(s[:-1])
    units = max(s[s.index(tens) + 1 :])
    return int(tens + units)


def Part1(input_stream: typing.TextIO) -> int:
    total_sum = 0
    for line in map(str.strip, input_stream):
        if not line:
            continue
        total_sum += reduce(line)
    return total_sum


if __name__ == "__main__":
    print(Part1(sys.stdin))
