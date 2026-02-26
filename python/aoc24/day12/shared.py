import typing

_coord: typing.TypeAlias = tuple[int, int]


class _Field:
    # Corrected to plural: __slots__
    __slots__ = ["width", "height", "data"]

    def __init__(
        self, width: int = 0, height: int = 0, data: list[bytearray] | None = None
    ):
        self.width = width
        self.height = height
        self.data = data if data is not None else []

    @classmethod
    def parse_input(cls, file_stream: typing.TextIO) -> "_Field":
        data: list[bytearray] = []
        width = 0

        for line in file_stream:
            clean_line = line.strip()

            if not clean_line:
                if data:  # terminate parsing after hitting empty line after content
                    break
                continue

            current_width = len(clean_line)
            if width == 0:
                width = current_width
            elif current_width != width:
                raise ValueError(
                    f"Irregular line length: expected {width}, got {current_width}"
                )

            data.append(bytearray(clean_line, "ascii"))

        return cls(width=width, height=len(data), data=data)

    def get(self, x: int, y: int) -> int:
        return self.data[y][x]

    def collect_islands(self) -> list[set[_coord]]:
        islands: list[set[_coord]] = list()
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        visited: set[_coord] = set()
        for x in range(self.width):
            for y in range(self.height):
                start = (y, x)
                if start in visited:
                    continue
                # flood fill cells into buffer
                buffer: set[_coord] = set()
                toVisit = [start]
                groupValue = self.data[y][x]
                while len(toVisit) > 0:
                    curPos = toVisit.pop()
                    visited.add(curPos)
                    buffer.add(curPos)
                    curY, curX = curPos
                    for dirY, dirX in dirs:
                        nextY, nextX = curY + dirY, curX + dirX
                        # out of bounds checks
                        if (
                            nextY < 0
                            or nextX < 0
                            or nextY >= self.height
                            or nextX >= self.width
                        ):
                            continue
                        nextPos = (nextY, nextX)
                        if nextPos in visited:
                            continue
                        if (
                            self.data[nextY][nextX] == groupValue
                            and nextPos not in buffer
                        ):
                            toVisit.append(nextPos)
                islands.append(buffer)
        return islands

    @staticmethod
    def count_edges(island: set[_coord], straight_edges: bool = False):
        """
        Counts the number of edges of the island
        """

        steps = 0
        dirChanges = 0

        edgeChecks = [
            (0, 1),
            (1, 1),
            (1, 0),
            (1, -1),
            (0, -1),
            (-1, -1),
            (-1, 0),
            (-1, 1),
        ]
        for y, x in island:
            for i in range(0, 8, 2):
                cellA = (y + edgeChecks[i][0], x + edgeChecks[i][1])
                cellB = (y + edgeChecks[(i + 2) % 8][0], x + edgeChecks[(i + 2) % 8][1])
                if cellA not in island:
                    steps += 1
                    if cellB not in island:
                        dirChanges += 1
                elif cellB in island:
                    cellC = (
                        y + edgeChecks[(i + 1) % 8][0],
                        x + edgeChecks[(i + 1) % 8][1],
                    )
                    if cellC not in island:
                        dirChanges += 1
        return steps, dirChanges
