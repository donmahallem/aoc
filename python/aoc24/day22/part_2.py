import functools
import codecs
import numpy as np

np.set_printoptions(linewidth=200)
GENERATIONS = 2001
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


def generatePricePattern(input_data, iterations=GENERATIONS):
    data_np = np.zeros((len(input_data), iterations, 2), dtype=np.int32)
    for i in range(len(input_data)):
        summe = input_data[i]
        data_np[i, 0, 0] = input_data[i] % 10
        for j in range(1, iterations):
            summe = calc(summe)
            data_np[i, j, 0] = summe % 10
            data_np[i, j, 1] = summe % 10 - data_np[i, j - 1, 0]
    return data_np


data_np = generatePricePattern(data, GENERATIONS)


def generatePatterns(data_np):
    output = set()
    vendor_pattern_dict = dict()
    for vendor_idx in range(data_np.shape[0]):
        vendor_idx = int(vendor_idx)
        if not (vendor_idx in vendor_pattern_dict):
            vendor_pattern_dict[vendor_idx] = dict()
        for j in range(0, data_np.shape[1] - 4):
            test = tuple(
                [
                    int(data_np[vendor_idx, j, 1]),
                    int(data_np[vendor_idx, j + 1, 1]),
                    int(data_np[vendor_idx, j + 2, 1]),
                    int(data_np[vendor_idx, j + 3, 1]),
                ]
            )
            if not (test in vendor_pattern_dict[vendor_idx]):
                vendor_pattern_dict[vendor_idx][test] = int(
                    data_np[vendor_idx, j + 3, 0]
                )
            output.add(test)
    return output, vendor_pattern_dict


unique_patterns, vendor_pattern_dict = generatePatterns(data_np)

from tqdm import tqdm

last_summe = None
last_pattern = None
for pattern_idx in tqdm(unique_patterns):
    summe = 0
    for vendor_idx in range(data_np.shape[0]):
        if not (pattern_idx in vendor_pattern_dict[vendor_idx]):
            continue
        summe += vendor_pattern_dict[vendor_idx][pattern_idx]
    if last_summe == None and last_pattern == None:
        last_pattern = pattern_idx
        last_summe = summe
    elif last_summe < summe:
        last_pattern = pattern_idx
        last_summe = summe

print(last_pattern)
print(last_summe)
