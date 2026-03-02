import sys
import typing

_WORDS: list[str] = [
    "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"
]


def _parse_line(line: str) -> int:
    first: int | None = None
    last: int | None = None

    for i in range(len(line)):
        val: int | None = None
        ch = line[i]
        if ch.isdigit():
            val = int(ch)
        else:
            for word_idx, word in enumerate(_WORDS):
                if line[i:i + len(word)] == word:
                    val = word_idx + 1
                    break
        if val is not None:
            if first is None:
                first = val
            last = val

    if first is None:
        return 0
    return first * 10 + last


def Part2(input: typing.TextIO) -> int:
    return sum(_parse_line(line.rstrip("\n")) for line in input)


if __name__ == "__main__":
    print(Part2(sys.stdin))
