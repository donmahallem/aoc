import typing
import math
from functools import cache

def Part1(input: typing.TextIO,blinks:int=25) -> int:
    data = [int(a) for a in input.readline().split(" ")]
    summe = 0
    for i in data:
        summe += christmas_tree(i, blinks)
    return summe

@cache
def christmas_tree(item, depth):
    if depth == 0:
        return 1
    if item == 0:
        result = christmas_tree(1, depth - 1)
        return result
    item_length = int(math.log10(item)) + 1
    if item_length % 2 == 0:
        split_barrier = 10 ** (item_length / 2)
        result = christmas_tree(int(item / split_barrier), depth - 1) + christmas_tree(
            int(item % split_barrier), depth - 1
        )
        return result
    else:
        result = christmas_tree(item * 2024, depth - 1)
        return result


