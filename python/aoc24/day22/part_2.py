import typing
from .shared import _parseInput


def Part2(input: typing.TextIO) -> int:
    mask = 16777215  # 0xFFFFFF
    kmask = 1048575  # 0xFFFFF

    # cache for the sums of the seen keys, and which seed_id was the last to update it
    psums = [0] * 1048576
    seen = [-1] * 1048576

    for seed_id, tmp in enumerate(_parseInput(input)):
        prev_val = tmp % 10
        key = 0

        # instead of checking every iteration if three have passed
        # just run them first and than just run the remaining iterations
        for _ in range(3):
            tmp = (tmp ^ (tmp << 6)) & mask
            tmp = (tmp ^ (tmp >> 5)) & mask
            tmp = (tmp ^ (tmp << 11)) & mask

            curr = tmp % 10
            key = ((key << 5) & kmask) | (curr - prev_val + 9)
            prev_val = curr

        # --- THE HOT LOOP ---
        for _ in range(
            1996
        ):  # 2000 total iterations, minus 1 for initial state, minus 3 for peeled. (1996)
            tmp = (tmp ^ (tmp << 6)) & mask
            tmp = (tmp ^ (tmp >> 5)) & mask
            tmp = (tmp ^ (tmp << 11)) & mask

            curr = tmp % 10
            key = ((key << 5) & kmask) | (curr - prev_val + 9)

            if seen[key] != seed_id:
                seen[key] = seed_id
                psums[key] += curr

            prev_val = curr

    return max(psums)
