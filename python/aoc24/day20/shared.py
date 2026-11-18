import array
import sys
import typing

_TRACK_BITS = str.maketrans("#.SE", "0111")
# Each packed value has 16 bits. Values stay below _BIAS, and bit 15 holds
# the result of a range check.
_FIELD_BITS = 16
_BIAS = 1 << 14
_TOP = 1 << 15


def _walk_track(lines: list[str]) -> list[int]:
    width = len(lines[0])
    stride = width + 1
    grid = "#".join(lines)
    pos = grid.index("S")
    end = grid.index("E")
    y, x = divmod(pos, stride)
    u, v = y + x, y - x + width
    moves = {
        1: (1, -1, stride, -stride),
        -1: (-1, 1, stride, -stride),
        stride: (1, 1, 1, -1),
        -stride: (-1, -1, 1, -1),
    }
    direction = next(d for d in (1, -1, stride, -stride) if grid[pos + d] != "#")
    du, dv, turn_a, turn_b = moves[direction]
    coords = [u, v]
    extend = coords.extend
    while pos != end:
        if grid[pos + direction] == "#":
            direction = turn_a if grid[pos + turn_a] != "#" else turn_b
            du, dv, turn_a, turn_b = moves[direction]
        pos += direction
        u += du
        v += dv
        extend((u, v))
    return coords


def _count_close_pairs(lines: list[str], max_distance: int) -> int:
    """Counts unordered track cell pairs with a distance of 1 to max_distance."""
    row_stride = len(lines[0]) + max_distance
    # The zero columns between rows stop horizontal shifts from wrapping.
    track = int(("0" * max_distance).join(lines).translate(_TRACK_BITS), 2)
    total = 0
    for dy in range(max_distance + 1):
        for dx in range(dy - max_distance, max_distance - dy + 1):
            if dy or dx > 0:
                total += (track & (track >> (dy * row_stride + dx))).bit_count()
    return total


def count_cheats(input: typing.TextIO, min_savings: int, max_distance: int) -> int:
    """Counts the cheats of at most max_distance steps that save min_savings.

    A cheat from path index i to j with distance d saves j - i - d steps.
    The result is all close pairs minus the close pairs that save too little.
    Only pairs with j - i < min_savings + max_distance can save too little.
    """
    lines = input.read().split()
    coords = _walk_track(lines)
    n = len(coords) // 2
    if n < min_savings + 3:
        return 0
    total = _count_close_pairs(lines, max_distance)

    max_gap = min(n, min_savings + max_distance)
    # Padding cells are far from all track cells, so the shifted values of the
    # last cells never match.
    coords.extend([_BIAS - 1] * (2 * max_gap))
    values = array.array("H", coords)
    if sys.byteorder == "big":
        values.byteswap()
    packed = int.from_bytes(values.tobytes(), "little")
    unit = int.from_bytes(b"\x01\x00" * len(coords), "little")
    base = unit * _BIAS - packed
    # Bit 15 of the u field of each real path cell.
    u_flags = int.from_bytes(b"\x00\x80\x00\x00" * n, "little")
    bounds: dict[int, tuple[int, int]] = {}

    def count_within(gap: int, limit: int) -> int:
        """Counts path indices i where cell i + gap has a distance <= limit."""
        bound = bounds.get(limit)
        if bound is None:
            bound = bounds[limit] = (
                base + unit * (_TOP - _BIAS + limit),
                base + unit * (_TOP - _BIAS - 1 - limit),
            )
        shifted = packed >> (2 * _FIELD_BITS * gap)
        # Each field holds _BIAS + delta. Bit 15 of the xor is set if
        # -limit <= delta <= limit.
        within = (shifted + bound[0]) ^ (shifted + bound[1])
        return (within & (within >> _FIELD_BITS) & u_flags).bit_count()

    too_short = 0
    for gap in range(1, max_gap):
        # The distance never exceeds the gap.
        too_short += n - gap if gap <= max_distance else count_within(gap, max_distance)
        if gap > min_savings:
            too_short -= count_within(gap, gap - min_savings)
    return total - too_short
