import typing


def _parseInput(input: typing.TextIO) -> list[int]:
    """Read all integers from the input file."""
    return [int(line) for line in input if line.strip()]


def _loop_steps(seed: int) -> int:
    mask = (1 << 24) - 1
    tmp = seed
    for _ in range(2000):
        tmp = (tmp ^ (tmp << 6)) & mask
        tmp = (tmp ^ (tmp >> 5)) & mask
        tmp = (tmp ^ (tmp << 11)) & mask
    return tmp
