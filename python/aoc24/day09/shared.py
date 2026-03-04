import typing

def _parse_input(input: typing.TextIO) -> list[int]:
    # single line of digits
    return [int(ch) for ch in input.readline()]
