import functools
import codecs
import numpy as np

np.set_printoptions(linewidth=200)
test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = [int(line.strip()) for line in f.readlines()]

@functools.cache
def calc(val):
    PRUNE_VALUE = 16777216
    out = (val ^ (val * 64)) % PRUNE_VALUE
    out = ((out // 32) ^ out) % PRUNE_VALUE
    out = ((out * 2048) ^ out) % PRUNE_VALUE
    return out

summe = 0
for initial_value in data:
    test = initial_value
    for i in range(2000):
        test = calc(test)
    summe += test
print(summe)
