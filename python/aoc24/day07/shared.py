import typing

Equation = tuple[int, tuple[int, ...]]

def _stream_rows(file_stream: typing.TextIO) -> typing.Iterator[Equation]:
    for line in file_stream:
        if not line.strip():
            continue

        target, numbers = line.split(":")
        yield int(target), tuple(int(x) for x in numbers.split())

def parseRows(file_stream: typing.TextIO) -> list[Equation]:
    data = []
    for line in file_stream:
        if not line.strip():
            continue

        target, numbers = line.split(":")
        data.append((int(target), tuple(int(x) for x in numbers.split())))

    return data
