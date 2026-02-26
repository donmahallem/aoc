import typing

type _splitter = set[int]

def __parseInput(
        input: typing.TextIO) -> tuple[tuple[int, int], _splitter, int, int]:
    splitter: _splitter = set()
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
                    splitter.add(i + height * width)
        height += 1
    return start, splitter, width, height
