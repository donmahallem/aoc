import typing
import sys
import re


def Part2(input_stream: typing.TextIO) -> int:
    data = input_stream.read()

    total = 0
    mul_pattern = re.compile(r"mul\((\d+),(\d+)\)")

    parts = data.split("don't()")

    for a, b in mul_pattern.findall(parts[0]):
        total += int(a) * int(b)

    for part in parts[1:]:
        if "do()" in part:
            sub_chunks = part.split("do()")
            for chunk in sub_chunks[1:]:
                for a, b in mul_pattern.findall(chunk):
                    total += int(a) * int(b)

    return total


if __name__ == "__main__":
    print(Part2(sys.stdin))
