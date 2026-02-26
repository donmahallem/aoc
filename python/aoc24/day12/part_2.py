import typing
import numpy as np

from .shared import _Field


def Part2(input: typing.TextIO) -> int:
    field = _Field.parse_input(input)
    islands = field.collect_islands()
    total = 0
    for island in islands:
        _, edges = _Field.count_edges(island)
        total += len(island) * edges
    return total
