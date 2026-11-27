import typing


def parse_input(input: typing.TextIO) -> list[int]:
    return list(map(int, input.read().strip().split(",")))
