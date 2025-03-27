import functools
import itertools
import codecs

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = [line.strip() for line in f.readlines()]


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


interconnected = set()
for key in connections.keys():
    test = connections[key]
    for test_key in test:
        for con3 in connections[test_key]:
            if con3 in connections[key]:
                items = tuple(sorted([key, test_key, con3]))
                interconnected.add(items)
print("Combinations", len(interconnected))


def countT(items):
    sum = 0
    for item in items:
        for a in item:
            if "t" == a[0]:
                sum += 1
                break
    return sum


print("Counted Ts", countT(interconnected))
