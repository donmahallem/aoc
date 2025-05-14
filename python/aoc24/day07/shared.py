import typing

Equitation = tuple[int, list[int]]


def parseRows(input: typing.TextIO) -> list[Equitation]:
    lines: list[str] = [a.strip() for a in input.readlines()]

    split_lines: list[list[str]] = [row.split(":") for row in lines]
    data: list[Equitation] = [(int(result[0].strip()),
                               [int(a) for a in result[1].strip().split(" ")])
                              for result in split_lines]
    return data
