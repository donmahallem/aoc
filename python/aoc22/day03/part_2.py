import sys
import typing
from .shared import _map_char


def Part2(input_data: typing.TextIO) -> int:
    current_value: int = 0
    lines = input_data.read().splitlines()
    for i in range(0, len(lines), 3):
        group = lines[i : i + 3]
        if len(group) < 3:
            continue
        common_chars = set(group[0]) & set(group[1]) & set(group[2])
        if common_chars:
            current_value += sum(_map_char(ord(c)) for c in common_chars)
    return current_value


if __name__ == "__main__":
    print(Part2(sys.stdin))
