import math
import typing
from functools import lru_cache
from .shared import parseRows


def Part2(input_file: typing.TextIO) -> int:
    rows = parseRows(input_file)
    total_sum = 0

    @lru_cache(maxsize=None)
    def can_solve(current_target: int, idx: int, terms_tuple: tuple[int, ...]) -> bool:
        if idx == 0:
            return terms_tuple[0] == current_target

        if current_target < 0:
            return False

        last_val = terms_tuple[idx]

        if current_target > last_val:
            if can_solve(current_target - last_val, idx - 1, terms_tuple):
                return True

        if current_target % last_val == 0:
            if can_solve(current_target // last_val, idx - 1, terms_tuple):
                return True

        if last_val < 10:
            mag = 10
        elif last_val < 100:
            mag = 100
        elif last_val < 1000:
            mag = 1000
        else:
            mag = 10 ** (int(math.log10(last_val)) + 1)

        if current_target % mag == last_val:
            if can_solve(current_target // mag, idx - 1, terms_tuple):
                return True

        return False

    for target, terms in rows:
        if can_solve(target, len(terms) - 1, terms):
            total_sum += target

        # Clear cache
        can_solve.cache_clear()

    return total_sum
