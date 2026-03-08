from util.typed_coord import Coord, CoordX, CoordY
import typing


def _parse_field(
    input_data: typing.TextIO,
    double_width: bool = False,
) -> tuple[set[Coord], set[Coord], Coord, list[Coord]]:
    walls: set[Coord] = set()
    player_position: Coord | None = None
    boxes: set[Coord] = set()
    currentRow = 0
    xMultiplier = 2 if double_width else 1
    for line in input_data:
        line = line.strip()
        if line == "":
            break
        for col in range(0, len(line)):
            match line[col]:
                case "#":
                    walls.add((currentRow, col * xMultiplier))
                    if double_width:
                        walls.add((currentRow, col * xMultiplier + 1))
                case "O":
                    boxes.add((currentRow, col * xMultiplier))
                case "@":
                    player_position = (currentRow, col * xMultiplier)
        currentRow += 1
    movements: list[Coord] = list()
    for line in input_data:
        line = line.strip()
        if line == "":
            break
        for dir in line:
            match dir:
                case "<":
                    movements.append((0, -1))
                case ">":
                    movements.append((0, 1))
                case "v":
                    movements.append((1, 0))
                case "^":
                    movements.append((-1, 0))

    if player_position is None:
        raise ValueError("Player position not found in input")
    return walls, boxes, player_position, movements
