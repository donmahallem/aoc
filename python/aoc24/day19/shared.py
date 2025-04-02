import typing

def parseInput(input: typing.TextIO) -> tuple[list[str], list[str]]:
    data = input.readlines()

    available_patterns = [line.strip() for line in data[0].strip().split(",")]
    needed_patterns = [line.strip() for line in data[2:]]

    return available_patterns, needed_patterns
