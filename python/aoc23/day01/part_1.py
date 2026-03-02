import sys
import typing


def Part1(input: typing.TextIO) -> int:
    total = 0
    for line in input:
        line = line.rstrip("\n")
        digits = [ch for ch in line if ch.isdigit()]
        if digits:
            total += int(digits[0]) * 10 + int(digits[-1])
    return total


if __name__ == "__main__":
    print(Part1(sys.stdin))
