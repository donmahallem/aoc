import typing
from .shared import _Field, _coord


def count_edges(coords: set[_coord]):
    edge_count = 0
    dirs = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    for coord in coords:
        y, x = coord
        for dir_y, dir_x in dirs:
            new_x, new_y = dir_x + x, dir_y + y
            if (new_y, new_x) in coords:
                continue
            edge_count += 1
    return edge_count


def Part1(input: typing.TextIO) -> int:
    field = _Field.parse_input(input)
    islands = field.collect_islands()

    total = 0
    for island in islands:
        num_edges, _ = _Field.count_edges(island)
        total += len(island) * num_edges
    return total
