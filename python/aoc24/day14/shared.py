import typing
import re
from util.typed_coord import CoordX, CoordY


def _parse_input(
    input: typing.TextIO,
) -> tuple[list[CoordX], list[CoordY], list[CoordX], list[CoordY]]:
    text = input.read()
    nums = list(map(int, re.findall(r"[-+]?\d+", text)))

    pxs: list[CoordX] = nums[0::4]  # type: ignore[assignment]
    pys: list[CoordY] = nums[1::4]  # type: ignore[assignment]
    vxs: list[CoordX] = nums[2::4]  # type: ignore[assignment]
    vys: list[CoordY] = nums[3::4]  # type: ignore[assignment]

    return pxs, pys, vxs, vys
