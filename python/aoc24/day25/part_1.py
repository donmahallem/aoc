import functools
import itertools
import codecs
import numpy as np

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt",
                 encoding="utf8") as f:
    data = f.read()
    locks = data.split("\r\n\r\n")

locks2 = []
keys2 = []
is_key = True
for lock in locks:
    lock_rows = lock.split("\r\n")
    translatetd_lock = [-1] * len(lock_rows[0])
    for i, lock_row in enumerate(lock_rows):
        if i == 0:
            is_key = lock_row == "....."
        for lock_column, lock_char in enumerate(lock_row):
            if lock_char == "#":
                translatetd_lock[lock_column] += 1
    if is_key:
        keys2.append(translatetd_lock)
    else:
        locks2.append(translatetd_lock)

locks_np = np.array(locks2, dtype=np.uint8)
keys_np = np.array(keys2, dtype=np.uint8)

matching_pairs = 0
for lock_num in range(locks_np.shape[0]):
    for key_num in range(keys_np.shape[0]):
        # print(locks_np[lock_num]+keys_np[key_num])
        if np.all((locks_np[lock_num] + keys_np[key_num]) <= [5, 5, 5, 5, 5]):
            matching_pairs += 1

print(matching_pairs)
