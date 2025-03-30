import sys
import typing


def Part2(input: typing.TextIO) -> int:
    l1 = []
    l2 = []
    data = input.readlines()
    for line in data:
        a = line.strip().split()
        if len(a) != 2:
            continue
        l1.append(int(a[0]))
        l2.append(int(a[1]))

    occurences = {key: len([1 for item2 in l2 if item2 == key]) for key in set(l2)}
    return sum(occurences[item] * item if item in occurences else 0 for item in l1)


if __name__ == "__main__":
    Part2(sys.stdin)
