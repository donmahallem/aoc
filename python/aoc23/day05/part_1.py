import sys
import typing
from .shared import _parse_input


def Part1(input: typing.TextIO) -> int:
    seeds, layers = _parse_input(input)

    lowest_value: int | None = None

    for seed in seeds:
        current_value = seed
        for options in layers:
            for targetStart, sourceStart, length in options:
                if sourceStart <= current_value < sourceStart + length:
                    current_value = current_value - sourceStart + targetStart
                    break

        if lowest_value is None or current_value < lowest_value:
            lowest_value = current_value
    if lowest_value is None:
        raise ValueError("No seeds provided")
    return lowest_value


if __name__ == "__main__":
    print(Part1(sys.stdin))
