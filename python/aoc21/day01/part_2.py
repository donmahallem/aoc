import re
import sys
import typing


def Part2(input_data: typing.TextIO) -> int:
    ring_buffer: list[int] = [0, 0, 0, 0]
    ring_index: int = 0
    increased: int = 0
    for line in input_data.read().splitlines():
        if not line:
            continue
        value = int(line)
        ring_buffer[ring_index % 4] = value
        if ring_index >= 3:
            prevSumIndex = (ring_index + 1) % 4
            currSumIndex = ring_index % 4

            # We compare (A+B+C)>(B+C+D) <=> A > D
            if ring_buffer[currSumIndex] > ring_buffer[prevSumIndex]:
                increased += 1
        ring_index += 1
    return increased


if __name__ == "__main__":
    print(Part2(sys.stdin))
