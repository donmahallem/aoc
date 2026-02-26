import sys
import typing


def get_max_subsequence(digits: typing.List[int], seqLength: int) -> typing.List[int]:
    n = len(digits)
    if seqLength >= n:
        return digits[:]  # nothing to remove
    res: typing.List[int] = []
    for i, d in enumerate(digits):
        # pop while we can and replacing yields larger lexicographic subsequence
        while res and res[-1] < d and len(res) - 1 + (n - i) >= seqLength:
            # check if we can pop
            # check if previous digit is smaller
            # check if enough digits remain to fill seqLength
            res.pop()
        if len(res) < seqLength:
            res.append(d)
    # trim in case
    return res[:seqLength]


def digits_to_int(seq: typing.List[int]) -> int:
    v = 0
    for d in seq:
        v = v * 10 + d
    return v


def Part2(input: typing.TextIO) -> int:
    total_sum = 0
    for line in input:
        line = line.strip()
        digits = [int(ch) for ch in line]
        best_seq = get_max_subsequence(digits, 12)
        total_sum += digits_to_int(best_seq)
    return total_sum


if __name__ == "__main__":
    print(Part2(sys.stdin))
