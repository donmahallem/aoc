import math


def _find_repeated_blocks(lbound: int, ubound: int) -> dict[int, int]:
    result: dict[int, int] = {}

    min_len = len(str(lbound))
    max_len = len(str(ubound))

    for total_length in range(min_len, max_len + 1):
        for block_len in range(1, total_length // 2 + 1):
            if total_length % block_len != 0:
                continue

            projected_block_size = total_length // block_len

            mult = (10**total_length - 1) // (10**block_len - 1)

            lower_constraint = (lbound + mult - 1) // mult
            upper_constraint = ubound // mult

            min_block_val = 10 ** (block_len - 1)
            max_block_val = 10**block_len - 1

            start_block = max(min_block_val, lower_constraint)
            end_block = min(max_block_val, upper_constraint)

            for block in range(start_block, end_block + 1):
                block_value = block * mult
                if 10 ** (total_length - 1) <= block_value < 10**total_length:
                    result[block_value] = projected_block_size

    return result
