import typing
from .part_1 import parseField, findInterconnected


def findLongest(
    connections: dict[str, list[str]], interconnected: set[tuple[str, str, str]]
):
    connecting = 0
    largerConnections: list[tuple[str, ...]] = list(interconnected)
    while connecting < len(largerConnections):
        connected = largerConnections[connecting]
        check = connected[-1]
        for to_check in connections[check]:
            if to_check in connected:
                continue
            counter = 0
            for i in range(0, len(connected) - 1):
                if to_check in connections[connected[i]]:
                    counter += 1
                else:
                    break
            if counter == len(connected) - 1:
                largerConnections[connecting] += (to_check,)
                break
            else:
                largerConnections[connecting] = largerConnections[connecting]
                connecting += 1
                break
    return largerConnections


def Part2(input: typing.TextIO) -> list[str]:
    connections = parseField(input)
    interconnected = findInterconnected(connections)
    max_groups = findLongest(connections, interconnected)
    max_groups = sorted(max_groups, key=len)
    return sorted(max_groups[-1], reverse=False)
