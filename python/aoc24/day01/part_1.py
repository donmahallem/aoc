import codecs
import sys
import typing

def Part1(input: typing.TextIO) -> int:
    l1 = []
    l2 = []
    data = input.readlines()
    for line in data:
        a = line.strip().split()
        if len(a) != 2:
            continue
        l1.append(int(a[0]))
        l2.append(int(a[1]))
    l1 = sorted(l1)
    l2 = sorted(l2)

    return sum(abs(a - b) for a, b in zip(l1, l2))

if __name__ == "__main__":
    Part1(sys.stdin)
