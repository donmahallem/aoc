import re
import sys
import typing

_FIRST = re.compile(r"[1-9]")
_LAST = re.compile(r".*([1-9])")


def Part1(input_data: typing.TextIO) -> int:
    total = 0
    first_search = _FIRST.search
    last_search = _LAST.search
    for line in input_data.read().splitlines():
        if not line:
            continue
        m1 = first_search(line)
        if m1:
            total += (ord(m1.group()) - 48) * 10 + (
                ord(last_search(line).group(1)) - 48  # type: ignore[union-attr]
            )
    return total


if __name__ == "__main__":
    print(Part1(sys.stdin))
