import typing


def _parse(input: typing.TextIO) -> list[list[int]]:
    return [list(map(int, line.split())) for line in input if line.strip()]
