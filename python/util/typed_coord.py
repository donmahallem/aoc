from typing import NewType

# Designating coord X
CoordX = NewType("CoordX", int)
# Designating coord Y
CoordY = NewType("CoordY", int)

# Designating a coordinate as a tuple of (Y, X)
Coord = NewType("Coord", tuple[CoordY, CoordX])
