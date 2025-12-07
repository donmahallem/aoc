import typing


def __parseInput(
        input: typing.TextIO) -> tuple[list[tuple[int, int]], list[str]]:
    start: tuple[int, int] = None
    splitter: dict[int, bool] = dict()
    width: int = -1
    height: int = 0
    for line in input:
        line = line.strip()
        if line == "":
            continue
        if width < 0:
            width = len(line)
        for i, c in enumerate(line):
            match c:
                case 'S':
                    start = (i, height)
                case '^':
                    splitter[i + height * width] = True
        height += 1
    return start, splitter, width, height
