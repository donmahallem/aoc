import typing
from .shared import _stream_rows


def Part2(input_file: typing.TextIO) -> int:
    total_sum = 0

    for target, terms in _stream_rows(input_file):
        n = len(terms)

        mags = [0] * n
        for i in range(n):
            v = terms[i]
            if v < 10:
                mags[i] = 10
            elif v < 100:
                mags[i] = 100
            elif v < 1000:
                mags[i] = 1000
            else:
                m = 10000
                while m <= v:
                    m *= 10
                mags[i] = m

        def solve(t: int, idx: int) -> bool:
            if idx == 0:
                return terms[0] == t
            if t <= 0:
                return False

            val = terms[idx]

            if t > val and solve(t - val, idx - 1):
                return True

            if t % val == 0 and solve(t // val, idx - 1):
                return True

            mag = mags[idx]
            if t % mag == val and solve(t // mag, idx - 1):
                return True

            return False

        if solve(target, n - 1):
            total_sum += target

    return total_sum
