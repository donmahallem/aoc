import sys
import typing
import math


def _count_options(time: int, distance: int) -> int:
    """Count integers t in (0, time) where t*(time-t) > distance."""
    if time <= 0:
        return 0
    disc = time * time - 4 * distance
    if disc < 0:
        return 0
    sqrt_disc = math.sqrt(disc)
    lower = (time - sqrt_disc) / 2
    upper = (time + sqrt_disc) / 2
    upper_int = int(upper)
    if upper == upper_int:
        upper_int -= 1
    return upper_int - int(lower)


def _parse_input(input: typing.TextIO) -> list[tuple[int, int]]:
    lines = [line.strip() for line in input if line.strip()]
    times = list(map(int, lines[0].split()[1:]))
    return list(zip(times, map(int, lines[1].split()[1:])))


def Part1(input: typing.TextIO) -> int:
    races = _parse_input(input)
    result = 1
    for time, distance in races:
        result *= _count_options(time, distance)
    return result


if __name__ == "__main__":
    print(Part1(sys.stdin))
