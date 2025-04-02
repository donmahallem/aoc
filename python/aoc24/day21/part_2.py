import typing
from .part_1 import calculate_puzzle_output

def Part2(input: typing.TextIO) -> int:
    data = [line.strip() for line in input.readlines()]

    return calculate_puzzle_output(data, 26)
