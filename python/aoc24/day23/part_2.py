import typing
from .part_1 import parseField, findInterconnected


def findLongest(connections, interconnected):
    connecting = 0
    interconnected = list(interconnected)
    while connecting < len(interconnected):
        connected = interconnected[connecting]
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
                interconnected[connecting] += (to_check,)
                break
            else:
                interconnected[connecting] = interconnected[connecting]
                connecting += 1
                break
    return interconnected


def Part2(input: typing.TextIO) -> int:
    connections = parseField(input)
    interconnected = findInterconnected(connections)
    max_groups = findLongest(connections, interconnected)
    max_groups = sorted(max_groups, key=len)
    return sorted(max_groups[-1], reverse=False)
