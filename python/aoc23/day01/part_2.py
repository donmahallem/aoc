import re
import sys
import typing

_DIGIT_MAP = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
}
_FIRST = re.compile(r"one|two|three|four|five|six|seven|eight|nine|[1-9]")
_LAST = re.compile(r".*(one|two|three|four|five|six|seven|eight|nine|[1-9])")


def Part2(input_data: typing.TextIO) -> int:
    total = 0
    first_search = _FIRST.search
    last_search = _LAST.search
    digit_map = _DIGIT_MAP
    for line in input_data.read().splitlines():
        if not line:
            continue
        m1 = first_search(line)
        if m1:
            m2 = last_search(line)
            if m2:
                total += digit_map[m1.group()] * 10 + digit_map[m2.group(1)]
    return total


if __name__ == "__main__":
    print(Part2(sys.stdin))
