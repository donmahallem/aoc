import typing

from .shared import _parse_field
from util.typed_coord import CoordY, CoordX, Coord


def Part1(input: typing.TextIO) -> int:
    walls, boxes, player_pos, movements = _parse_field(input)

    for move in movements:
        dy, dx = move
        ny, nx = player_pos[0] + dy, player_pos[1] + dx
        next_pos = (ny, nx)

        if next_pos in walls:
            continue

        if next_pos not in boxes:
            # just move
            player_pos = next_pos
            continue

        ey, ex = ny + dy, nx + dx
        while (ey, ex) in boxes:
            ey, ex = ey + dy, ex + dx

        if (ey, ex) in walls:
            # Can't push
            continue

        boxes.discard(next_pos)
        boxes.add((ey, ex))
        player_pos = next_pos

    return sum(y * 100 + x for y, x in boxes)
