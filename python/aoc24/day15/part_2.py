import typing
from .shared import _parse_field
from util.typed_coord import CoordY, CoordX, Coord


def Part2(input: typing.TextIO) -> int:
    walls, boxes, player_pos, movements = _parse_field(input, double_width=True)

    for dy, dx in movements:
        ny: CoordY = player_pos[0] + dy
        nx: CoordX = player_pos[1] + dx
        nextPos: Coord = (ny, nx)
        # 1. Check if the player directly hits a wall
        if nextPos in walls:
            continue

        # 2. Check if the player hits a box (checking left and right sides of boxes)
        hit_box: Coord | None = None
        if nextPos in boxes:
            hit_box = nextPos
        elif (ny, nx - 1) in boxes:
            hit_box = (ny, nx - 1)  # type: ignore[assignment]

        # If it's an empty space, just move the player
        if hit_box is None:
            player_pos = nextPos
            continue

        # 3. BFS to find all boxes that need to move in a chain reaction
        boxes_to_move = {hit_box}
        queue = [hit_box]
        blocked = False

        while queue:
            by, bx = queue.pop(0)
            nby: CoordY = by + dy
            nbx: CoordX = bx + dx

            # A wide box covers nbx and nbx + 1. Check if either side hits a wall.
            if (nby, nbx) in walls or (nby, nbx + 1) in walls:
                blocked = True
                break

            # Check if this moving box bumps into other boxes
            for py, px in [(nby, nbx), (nby, nbx + 1)]:
                b: Coord | None = None
                if (py, px) in boxes:
                    b = (py, px)
                elif (py, px - 1) in boxes:
                    b = (py, px - 1)

                if b and b not in boxes_to_move:
                    boxes_to_move.add(b)
                    queue.append(b)

        # 4. Execute the move if nothing hit a wall
        if not blocked:
            # Remove all old positions first to prevent self-collision
            boxes.difference_update(boxes_to_move)
            # Add all new positions
            for by, bx in boxes_to_move:
                boxes.add((by + dy, bx + dx))

            player_pos = (ny, nx)

    # Calculate final GPS coordinates
    return sum(y * 100 + x for y, x in boxes)
