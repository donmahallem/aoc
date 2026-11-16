from operator import le
import re
import sys
import typing
import math
from .shared import _map_char

_OFFSET_a = ord("a")
_OFFSET_z = ord("z")
_OFFSET_A = ord("A")
_OFFSET_Z = ord("Z")


def Part1(input_data: typing.TextIO) -> int:
    current_value: int = 0
    for line in input_data.read().splitlines():
        split_index = len(line) // 2
        l = line[:split_index]
        for i in range(split_index, len(line)):
            if line[i] in l:
                current_value += _map_char(ord(line[i]))
                break
    return current_value


if __name__ == "__main__":
    print(Part1(sys.stdin))
