import typing

_PRUNE_MASK: int = (1 << 24) - 1
_ITERATIONS: int = 2000


def _parseInput(input: typing.TextIO) -> list[int]:
    """Read all integers from the input file."""
    return [int(line.strip()) for line in input if line.strip()]


def _step(val: int) -> int:
    mask = _PRUNE_MASK
    tmp = (val ^ (val << 6)) & mask
    tmp = (tmp ^ (tmp >> 5)) & mask
    tmp = (tmp ^ (tmp << 11)) & mask
    return tmp
