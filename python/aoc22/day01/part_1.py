import sys
import typing


def Part1(input_data: typing.TextIO) -> int:
    current_value: int = 0
    max_value: int = 0
    for line in input_data.read().splitlines():
        if not line:
            if current_value > 0:
                max_value = max(max_value, current_value)
                current_value = 0
            continue
        current_value += int(line)
    return max(max_value, current_value)


if __name__ == "__main__":
    print(Part1(sys.stdin))
