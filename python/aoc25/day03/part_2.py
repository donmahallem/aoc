import sys


def reduce(s: str, k: int = 12) -> int:
    n = len(s)
    if n <= k:
        return int(s)
    start = 0
    out = []
    for remaining in range(k, 0, -1):
        window = s[start : n - remaining + 1]
        d = max(window)
        start += window.index(d) + 1
        out.append(d)
    return int("".join(out))


def Part2(input_stream) -> int:
    summe = 0
    for line in map(str.strip, input_stream):
        if not line:
            continue
        summe += reduce(line)

    return summe


if __name__ == "__main__":
    print(Part2(sys.stdin))
