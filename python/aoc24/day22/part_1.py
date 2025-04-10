import functools
import typing
import numpy as np

np.set_printoptions(linewidth=200)

def parseInput(input: typing.TextIO):
    return [int(line.strip()) for line in input.readlines()]


@functools.cache
def calc(val):
    PRUNE_VALUE = 16777216
    out = (val ^ (val * 64)) % PRUNE_VALUE
    out = ((out // 32) ^ out) % PRUNE_VALUE
    out = ((out * 2048) ^ out) % PRUNE_VALUE
    return out


def Part1(input: typing.TextIO) -> int:
    data = parseInput(input)
    summe = 0
    for initial_value in data:
        test = initial_value
        for i in range(2000):
            test = calc(test)
        summe += test
    return summe