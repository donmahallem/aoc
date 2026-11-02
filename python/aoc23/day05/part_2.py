import sys
import typing
from .shared import _parse_input


def _merge_ranges(ranges: list[tuple[int, int]]) -> list[tuple[int, int]]:
    if not ranges:
        return []
    # Sort ranges by start
    ranges.sort(key=lambda x: x[0])
    merged = [ranges[0]]
    for current in ranges[1:]:
        last_merged = merged[-1]
        if current[0] <= last_merged[1]:  # Overlapping or adjacent
            merged[-1] = (last_merged[0], max(last_merged[1], current[1]))  # Merge
        else:
            merged.append(current)
    return merged


def Part2(input: typing.TextIO) -> int:
    seeds, layers = _parse_input(input)

    # tracks the ranges within a "layer"
    ranges: list[tuple[int, int]] = []
    for idx in range(0, len(seeds), 2):
        ranges.append((seeds[idx], seeds[idx] + seeds[idx + 1] - 1))

    for options in layers:
        # collect the next set of ranges after transformation
        next_ranges: list[tuple[int, int]] = []
        while ranges:
            r_start, r_end = ranges.pop()
            matched = False
            for targetStart, sourceStart, length in options:
                s_start = sourceStart
                s_end = sourceStart + length - 1

                # Check for overlap
                o_start = max(r_start, s_start)
                o_end = min(r_end, s_end)
                if o_start <= o_end:
                    # Map the overlapping part
                    next_ranges.append(
                        (
                            targetStart + (o_start - s_start),
                            targetStart + (o_end - s_start),
                        )
                    )
                    # Queue remaining parts to be processed
                    if r_start < o_start:
                        ranges.append((r_start, o_start - 1))
                    if o_end < r_end:
                        ranges.append((o_end + 1, r_end))
                    matched = True
                    break

            if not matched:
                next_ranges.append((r_start, r_end))

        ranges = _merge_ranges(next_ranges)

    return min(r[0] for r in ranges)


if __name__ == "__main__":
    print(Part2(sys.stdin))
