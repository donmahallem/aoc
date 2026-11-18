import sys
import typing

from .shared import row_generator


def get_max_subsequence(digits: list[int], seq_length: int) -> list[int]:
    n = len(digits)
    if seq_length >= n:
        return digits[:]  # nothing to remove

    res: list[int] = []
    for i, d in enumerate(digits):
        # Pop elements smaller than current `d` as long as we have
        # enough remaining digits to still reach `seq_length`.
        while res and res[-1] < d and (len(res) - 1 + (n - i) >= seq_length):
            res.pop()
        if len(res) < seq_length:
            res.append(d)

    return res


def digits_to_int(seq: list[int]) -> int:
    v = 0
    for d in seq:
        v = v * 10 + d
    return v


def Part2(input_stream: typing.TextIO) -> int:
    total_sum = 0

    for line in row_generator(input_stream):
        digits = line
        best_seq = get_max_subsequence(digits, 12)
        total_sum += digits_to_int(best_seq)

    return total_sum


if __name__ == "__main__":
    print(Part2(sys.stdin))
