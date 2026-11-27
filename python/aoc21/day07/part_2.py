import functools
import sys
import typing

from .shared import parse_input


def Part2(input_data: typing.TextIO) -> int:
    data = parse_input(input_data)
    sorted_data = sorted(data)
    left_idx = min(sorted_data)
    right_idx = max(sorted_data)

    def gauss(n: int) -> int:
        return n * (n + 1) // 2

    @functools.cache
    def calculate_cost(target: int) -> int:
        return sum(gauss(abs(x - target)) for x in sorted_data)

    while True:
        # Calculate the gradient of the current middle point
        # If the distance is increasing in both directions, we have found the minimum
        mid_idx = (left_idx + right_idx) // 2
        mid_cost = calculate_cost(mid_idx)
        left_cost = calculate_cost(mid_idx - 1)
        right_cost = calculate_cost(mid_idx + 1)
        if mid_cost <= left_cost and mid_cost <= right_cost:
            return mid_cost
        if left_cost < right_cost:
            right_idx = mid_idx
        elif right_cost < left_cost:
            left_idx = mid_idx


if __name__ == "__main__":
    print(Part2(sys.stdin))
