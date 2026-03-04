import typing
from .shared import _parse_input


def Part1(input: typing.TextIO) -> int:
    data: list[int] = _parse_input(input)

    pos = 0
    total = 0
    idxL = 0
    idxR = len(data) - 1
    if idxR % 2 != 0:
        idxR -= 1

    r_file_rem = data[idxR]

    while idxL < idxR:
        size = data[idxL]
        if size > 0:
            f_id = idxL >> 1
            total += (f_id * size * (2 * pos + size - 1)) // 2
            pos += size
        idxL += 1

        if idxL >= idxR:
            break

        space = data[idxL]
        while space > 0 and idxL < idxR:
            if r_file_rem == 0:
                idxR -= 2
                if idxR <= idxL:
                    break
                r_file_rem = data[idxR]
                continue

            take = space if space < r_file_rem else r_file_rem
            f_id = idxR >> 1
            total += (f_id * take * (2 * pos + take - 1)) // 2

            pos += take
            space -= take
            r_file_rem -= take

        idxL += 1

    if r_file_rem > 0:
        total += ((idxR >> 1) * r_file_rem * (2 * pos + r_file_rem - 1)) // 2

    return total
