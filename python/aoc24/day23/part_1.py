import typing


def parseField(input: typing.TextIO) -> dict[str, list[str]]:
    data = [line.strip() for line in input.readlines()]

    connections: dict[str, list[str]] = dict()
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


from typing import TypeAlias

Triple: TypeAlias = tuple[str, str, str]


def findInterconnected(connections: dict[str, list[str]]) -> set[Triple]:
    conn_sets: dict[str, set[str]] = {
        k: set(v)
        for k, v in connections.items()
    }

    interconnected: set[Triple] = set()
    keys = sorted(connections.keys())

    for i, a in enumerate(keys):
        for j in range(i + 1, len(keys)):
            b = keys[j]
            if b in conn_sets[a]:
                common = conn_sets[a] & conn_sets[b]
                for c in common:
                    if c > b:
                        interconnected.add((a, b, c))

    return interconnected


def countT(items: set[tuple[str, str, str]]) -> int:
    sum: int = 0
    for item in items:
        for a in item:
            if "t" == a[0]:
                sum += 1
                break
    return sum


def Part1(input: typing.TextIO) -> int:
    data = parseField(input)
    interconnected = findInterconnected(data)
    return countT(interconnected)
