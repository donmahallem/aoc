import codecs
import numpy as np
from functools import cache

test_data = False

with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = f.readlines()

available_patterns=[ line.strip() for line in data[0].strip().split(",")]
needed_patterns=[line.strip() for line in data[2:]]

@cache
def checkPattern2(pattern):
    for test_pattern in available_patterns:
        if test_pattern==pattern:
            return True
        elif pattern[0:len(test_pattern)]==test_pattern:
            if checkPattern2(pattern[len(test_pattern):]):
                return True
    return False
results=0
for p in needed_patterns:
    if checkPattern2(p):
        results+=1

print(results)