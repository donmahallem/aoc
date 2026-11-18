import typing

_TO_BITS = str.maketrans("#.", "10")
_WIDTH = 5
_HEIGHT = 7
_ROW = _WIDTH + 1
# One block has 7 rows of 6 characters, plus one blank line.
_BLOCK = _HEIGHT * _ROW + 1


def _bits(cells: str) -> int:
    """Converts a string of "#" and "." cells into an integer, "#" as 1."""
    return int(cells.translate(_TO_BITS), 2)


def Part1(input: typing.TextIO) -> int:
    text = input.read().replace("\r", "").strip() + "\n\n"
    if len(text) % _BLOCK or set(text[_BLOCK - 1 :: _BLOCK]) != {"\n"}:
        raise ValueError("Expected 7x5 patterns separated by one blank line")

    # Each schematic has one bit. A set bit marks a key.
    key_mask = ~_bits(text[::_BLOCK]) & ((1 << (len(text) // _BLOCK)) - 1)
    # fits[c][h] holds the keys with a height of h or less in column c.
    # That is true if the cell in row 5 - h, column c is ".".
    fits = [
        [~_bits(text[(5 - h) * _ROW + col :: _BLOCK]) & key_mask for h in range(6)]
        for col in range(_WIDTH)
    ]
    fits_0, fits_1, fits_2, fits_3, fits_4 = fits

    total = 0
    for start in range(0, len(text), _BLOCK):
        if text[start] != "#":
            continue
        block = text[start : start + _BLOCK - 1]
        a, b, c, d, e = (5 - block[col::_ROW].count("#") + 1 for col in range(_WIDTH))
        total += (fits_0[a] & fits_1[b] & fits_2[c] & fits_3[d] & fits_4[e]).bit_count()
    return total
