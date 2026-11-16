import typing


def _map_char(c: str) -> int:
    match c:
        case "A" | "X":
            return 0
        case "B" | "Y":
            return 1
        case "C" | "Z":
            return 2
        case _:
            raise ValueError(f"Invalid character: {c}")


# generator for parsing
def _parser(input_data: typing.TextIO):
    for line in input_data.read().splitlines():
        if not line:
            continue
        if len(line) < 3:
            raise ValueError("Line is too short")
        yield _map_char(line[0]), _map_char(line[2])
