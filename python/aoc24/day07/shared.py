import typing


def parseRows(input: typing.TextIO) -> tuple[int, list[int]]:
    data = [a.strip() for a in input.readlines()]

    data = [row.split(":") for row in data]
    data = [
        (int(result[0].strip()), [int(a) for a in result[1].strip().split(" ")])
        for result in data
    ]
    return data
