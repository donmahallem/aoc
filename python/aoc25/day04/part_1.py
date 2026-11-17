import sys
import typing

_TO_BITS = str.maketrans("@.", "10")


def parseInput(input: typing.TextIO) -> tuple[int, int]:
    """Parses the field into one integer with one bit per cell.

    Returns the grid bits and the row stride. Each row has an empty pad bit,
    so horizontal neighbour shifts never wrap into the next row.
    """
    lines = input.read().split()
    stride = len(lines[0]) + 1
    return int("0".join(lines).translate(_TO_BITS), 2), stride


def accessible(grid: int, stride: int) -> int:
    """Returns the bits of all rolls with fewer than four neighbour rolls.

    The 3x3 sum includes the roll itself, so a roll stays if the sum is 5 or more.
    """
    # Sum each cell with its left and right cell: bits hs (1) and hc (2).
    left = grid << 1
    right = grid >> 1
    left_right = left ^ right
    hs = left_right ^ grid
    hc = (left & right) | (grid & left_right)
    # Sum the hs bits of three rows: bits s0 (1) and c0 (2).
    up = hs << stride
    down = hs >> stride
    up_down = up ^ down
    s0 = up_down ^ hs
    c0 = (up & down) | (hs & up_down)
    # Sum the hc bits of three rows: bits t1 (2) and t2 (4).
    up = hc << stride
    down = hc >> stride
    up_down = up ^ down
    t1 = up_down ^ hc
    t2 = (up & down) | (hc & up_down)
    # Add c0 and t1: bits u1 (2) and u2 (4).
    u1 = c0 ^ t1
    u2 = c0 & t1
    # The sum is s0 + 2 * u1 + 4 * (u2 + t2).
    blocked = (t2 & u2) | ((t2 | u2) & (u1 | s0))
    return grid & ~blocked


def Part1(input: typing.TextIO) -> int:
    grid, stride = parseInput(input)
    return accessible(grid, stride).bit_count()


if __name__ == "__main__":
    print(Part1(sys.stdin))
