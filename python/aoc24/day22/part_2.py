import typing
from .shared import _step, _ITERATIONS, _parseInput


def Part2(input: typing.TextIO) -> int:
    """
    Takes a preallocated buffer of size 2**20 to store all potential 5 digit sequences
    (there are 10**5 = 100000 possible sequences, but we can encode them in 20 bits).
    For each seed, we generate the sequence of digits and keep track of the sum of each 5 digit sequence we see.
    We also keep track of which seeds have contributed to each sequence to avoid double-counting.
    We return the maximum sum of any 5 digit sequence.

    """
    seeds = _parseInput(input)

    MASK = 0xFFFFF

    psums = [0] * (MASK + 1)
    seen = [-1] * (MASK + 1)

    for seed_id, seed in enumerate(seeds):
        tmp = seed
        prev_val = seed % 10
        key = 0

        for i in range(1, _ITERATIONS):
            tmp = _step(tmp)
            curr = tmp % 10

            diff = curr - prev_val + 9

            key = ((key << 5) & MASK) | diff

            if i >= 4:
                if seen[key] != seed_id:
                    seen[key] = seed_id
                    psums[key] += curr

            prev_val = curr

    return max(psums)
