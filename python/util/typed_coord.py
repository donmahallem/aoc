from typing import NewType

CoordX = NewType("CoordX", int)
CoordY = NewType("CoordY", int)

Coord = NewType("Coord", tuple[CoordX, CoordY])
