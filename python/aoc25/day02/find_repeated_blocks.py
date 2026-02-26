def find_repeated_blocks(lbound: int, ubound: int) -> dict[int, int]:
    result: dict[int, int] = {}
    min_len = len(str(lbound))
    max_len = len(str(ubound))

    for total_length in range(min_len, max_len + 1):
        for block_len in range(1, total_length // 2 + 1):
            if total_length % block_len != 0:
                continue
            k = total_length // block_len
            if k < 2:
                continue
            min_block = 10 ** (block_len - 1)
            max_block = 10**block_len - 1
            for block in range(min_block, max_block + 1):
                # No leading zeros
                if block_len > 1 and block < 10 ** (block_len - 1):
                    continue
                # Build repeated number
                mult = (10 ** (block_len * k) - 1) // (10**block_len - 1)
                v = block * mult
                if v < lbound or v > ubound:
                    continue
                # Ensure v has exactly total_length digits
                if not (10 ** (total_length - 1) <= v < 10**total_length):
                    continue
                result[v] = k
    return result
