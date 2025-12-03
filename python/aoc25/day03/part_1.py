import sys
import typing


def reduce(a: typing.List[int]) -> typing.List[int]:
    result = []
    prev = a[0]
    count = 1
    for ch in a[1:]:
        if ch == prev:
            count += 1
        else:
            if count >= 3:
                result.append(prev)
                result.append(prev)
            else:
                # append 1 or 2 copies
                result.extend([prev] * count)
            prev = ch
            count = 1
    # handle last run
    if count >= 3:
        result.append(prev)
        result.append(prev)
    else:
        result.extend([prev] * count)
    return result


def Part1(input: typing.TextIO) -> int:
    total_sum = 0
    for line in input:
        digits = [int(ch) for ch in line.strip()]
        digits = reduce(digits)
        max_val = 0
        for idxA in range(len(digits)):
            for idxB in range(idxA + 1, len(digits)):
                cur = digits[idxA] * 10 + digits[idxB]
                if cur > max_val:
                    max_val = cur
        total_sum += max_val
    return total_sum


if __name__ == "__main__":
    print(Part1(sys.stdin))
