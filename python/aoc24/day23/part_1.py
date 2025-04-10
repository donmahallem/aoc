import typing

def parseField(input: typing.TextIO) ->dict[str,list[str]]:
    data = [line.strip() for line in input.readlines()]

    connections = dict()
    for line in data:
        a, b = line.split("-")
        if a in connections:
            connections[a].append(b)
        else:
            connections[a] = list([b])
        if b in connections:
            connections[b].append(a)
        else:
            connections[b] = list([a])

    return connections

def findInterconnected(connections):
    interconnected = set()
    for key in connections.keys():
        test = connections[key]
        for test_key in test:
            for con3 in connections[test_key]:
                if con3 in connections[key]:
                    items = tuple(sorted([key, test_key, con3]))
                    interconnected.add(items)
    return interconnected

def countT(items) -> set[tuple[str]]:
    sum = 0
    for item in items:
        for a in item:
            if "t" == a[0]:
                sum += 1
                break
    return sum

def Part1(input: typing.TextIO) -> int:
    data = parseField(input)
    interconnected=findInterconnected(data)
    return countT(interconnected)