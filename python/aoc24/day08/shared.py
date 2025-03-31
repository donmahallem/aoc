import typing

def parseInput(input: typing.TextIO) -> tuple[typing.Dict[str, tuple[int,int]],int,int]:
    data = [a.strip() for a in input.readlines()]

    height = len(data)
    width = len(data[0])

    occurences = dict()
    for y in range(len(data)):
        for x in range(len(data[0])):
            if data[y][x] == ".":
                continue
            if data[y][x] in occurences:
                occurences[data[y][x]].append((y, x))
            else:
                occurences[data[y][x]] = list([(y, x)])
    return occurences,width,height

def inside(y, x,width,height):
    return x >= 0 and y >= 0 and x < width and y < height