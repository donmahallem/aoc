import codecs
import math

with codecs.open("data.txt", encoding="utf8") as f:
    data = [[int(num) for num in a.strip().split()] for a in f.readlines()]

data = data[0]
cache = dict()


def christmas_tree(item, depth):
    if depth == 0:
        return 1
    elif (item, depth) in cache:
        return cache[((item, depth))]
    if item == 0:
        result = christmas_tree(1, depth - 1)
        cache[(item, depth)] = result
        return result
    item_length = int(math.log10(item)) + 1
    if item_length % 2 == 0:
        split_barrier = 10 ** (item_length / 2)
        result = christmas_tree(int(item / split_barrier), depth - 1) + christmas_tree(
            int(item % split_barrier), depth - 1
        )
        cache[(item, depth)] = result
        return result
    else:
        result = christmas_tree(item * 2024, depth - 1)
        cache[(item, depth)] = result
        return result


blink_num = 25
sum = 0
for i in data:
    sum += christmas_tree(i, blink_num)
print(sum)
