import typing
from .shared import _stream_rows


def _is_valid(target: int, values: tuple[int, ...]) -> bool:
    def solve(t: int, idx: int) -> bool:
        if idx == 0:
            return values[0] == t
        if t <= 0:
            return False

        val = values[idx]

        if t > val and solve(t - val, idx - 1):
            return True

        if t % val == 0 and solve(t // val, idx - 1):
            return True

        return False

    return solve(target, len(values) - 1)


def Part1(file_stream: typing.TextIO) -> int:
    total_sum = 0

    for expected, values in _stream_rows(file_stream):
        if _is_valid(expected, values):
            total_sum += expected

    return total_sum
