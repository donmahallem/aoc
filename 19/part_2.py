import codecs
import numpy as np
from functools import cache

test_data = False

with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.readlines()

available_patterns=[ line.strip() for line in data[0].strip().split(",")]
needed_patterns=[line.strip() for line in data[2:]]

@cache
def checkPattern(pattern):
    #print(depth,pattern)
    valid_patterns=0
    for test_pattern in available_patterns:
        if test_pattern==pattern:
            valid_patterns+=1
        elif pattern[0:len(test_pattern)]==test_pattern:
            valid_patterns+=checkPattern(pattern[len(test_pattern):])
    return valid_patterns
results=0

summe=0
for p in needed_patterns:
    summe+=checkPattern(p)
print(summe)