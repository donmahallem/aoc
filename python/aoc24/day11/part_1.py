import typing
import math

def Part1(input: typing.TextIO,blinks:int=25) -> int:
    data = [int(a) for a in input.readline().split(" ")]
    cache = dict()
    summe = 0
    for i in data:
        summe += christmas_tree(i, blinks,cache)
    return summe


def christmas_tree(item, depth,cache):
    if depth == 0:
        return 1
    elif (item, depth) in cache:
        return cache[((item, depth))]
    if item == 0:
        result = christmas_tree(1, depth - 1,cache)
        cache[(item, depth)] = result
        return result
    item_length = int(math.log10(item)) + 1
    if item_length % 2 == 0:
        split_barrier = 10 ** (item_length / 2)
        result = christmas_tree(int(item / split_barrier), depth - 1,cache) + christmas_tree(
            int(item % split_barrier), depth - 1,cache
        )
        cache[(item, depth)] = result
        return result
    else:
        result = christmas_tree(item * 2024, depth - 1,cache)
        cache[(item, depth)] = result
        return result


