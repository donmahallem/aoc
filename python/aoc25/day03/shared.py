import typing


def row_generator(
    input_stream: typing.TextIO,
) -> typing.Generator[list[int], None, None]:
    offset_char = ord("0")
    for line in input_stream:
        line = line.strip()
        if not line:
            continue
        yield [ord(ch) - offset_char for ch in line]
