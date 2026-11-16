_OFFSET_a = ord("a")
_OFFSET_z = ord("z")
_OFFSET_A = ord("A")
_OFFSET_Z = ord("Z")


def _map_char(chr_ord: int) -> int:
    if _OFFSET_a <= chr_ord <= _OFFSET_z:
        return chr_ord - _OFFSET_a + 1
    elif _OFFSET_A <= chr_ord <= _OFFSET_Z:
        return chr_ord - _OFFSET_A + 27
    else:
        raise ValueError(f"Invalid character: {chr(chr_ord)}")
