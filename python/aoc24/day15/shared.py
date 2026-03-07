from util.typed_coord import Coord, CoordX, CoordY
import typing


def _parse_field(
    input_data: typing.TextIO,
) -> tuple[set[Coord], set[Coord], Coord, list[Coord]]:
    walls: set[Coord] = set()
    player_position: Coord | None = None
    boxes: set[Coord] = set()
    currentRow = 0
    for line in input_data:
        line = line.strip()
        if line == "":
            break
        for col in range(0, len(line)):
            match line[col]:
                case "#":
                    walls.add(Coord((CoordX(currentRow), CoordY(col))))
                case "O":
                    boxes.add(Coord((CoordX(currentRow), CoordY(col))))
                case "@":
                    player_position = Coord((CoordX(currentRow), CoordY(col)))
        currentRow += 1
    movements: list[Coord] = list()
    for line in input_data:
        line = line.strip()
        if line == "":
            break
        for dir in line:
            match dir:
                case "<":
                    movements.append(Coord((CoordX(0), CoordY(-1))))
                case ">":
                    movements.append(Coord((CoordX(0), CoordY(1))))
                case "v":
                    movements.append(Coord((CoordX(1), CoordY(0))))
                case "^":
                    movements.append(Coord((CoordX(-1), CoordY(0))))
    return walls, boxes, player_position, movements

def translateMovement(a) -> Coord:
    if a == "<":
        return (0, -1)
    elif a == ">":
        return (0, 1)
    elif a == "v":
        return (1, 0)
    elif a == "^":
        return (-1, 0)


CELL_BOX_LEFT = 4
CELL_BOX_RIGHT = 3
CELL_WALL = 2
CELL_BOX = 1
CELL_EMPTY = 0
