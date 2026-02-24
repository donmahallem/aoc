import typing

# directions encoded 0=up,1=right,2=down,3=left
_DY = (-1, 0, 1, 0)
_DX = (0, 1, 0, -1)

State = tuple[int, int, int]  # y, x, dir


def parseField(input: typing.TextIO) -> tuple[list[list[int]], State | None]:
    """Read guard map; walls = 1, open = 0.  Return grid and initial state."""
    data = [line.rstrip("\n") for line in input]
    rows = len(data)
    cols = len(data[0]) if rows else 0
    grid: list[list[int]] = [[0] * cols for _ in range(rows)]
    start: State | None = None
    for y, line in enumerate(data):
        for x, ch in enumerate(line):
            if ch == "#":
                grid[y][x] = 1
            elif ch == "^":
                start = (y, x, 0)
            elif ch == ">":
                start = (y, x, 1)
            elif ch == "v":
                start = (y, x, 2)
            elif ch == "<":
                start = (y, x, 3)
    return grid, start


def _move_next(grid: list[list[int]], state: State) -> State | None:
    y, x, d = state
    while True:
        ny = y + _DY[d]
        nx = x + _DX[d]
        # outside
        if not (0 <= ny < len(grid) and 0 <= nx < len(grid[0])):
            return None
        if grid[ny][nx] == 1:
            d = (d + 1) & 3
            continue
        return (ny, nx, d)


def Part1(input: typing.TextIO) -> int:
    grid, state = parseField(input)
    if state is None:
        raise ValueError("No guard found")
    visited: set[tuple[int, int]] = {(state[0], state[1])}
    while True:
        state = _move_next(grid, state)
        if state is None:
            break
        visited.add((state[0], state[1]))
    return len(visited)
