import typing
from collections import deque

Position = tuple[int, int]


def parseField(
    input: typing.TextIO,
) -> tuple[list[list[int]], Position | None, Position | None]:
    lines = [line.rstrip("\n") for line in input]
    if not lines:
        return [], None, None
    h = len(lines)
    w = len(lines[0])
    grid: list[list[int]] = [[0] * w for _ in range(h)]
    start: Position | None = None
    end: Position | None = None
    for y, line in enumerate(lines):
        for x, ch in enumerate(line):
            if ch == "#":
                grid[y][x] = 1
            elif ch == "S":
                start = (y, x)
            elif ch == "E":
                end = (y, x)
    return grid, start, end


def compute_shortest_path(
    grid: list[list[int]], start: Position, end: Position
) -> list[Position]:
    h = len(grid)
    w = len(grid[0]) if h else 0
    dq = deque([start])
    parent: dict[Position, Position | None] = {start: None}
    dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    while dq:
        y, x = dq.popleft()
        if (y, x) == end:
            break
        for dy, dx in dirs:
            ny, nx = y + dy, x + dx
            if 0 <= ny < h and 0 <= nx < w and grid[ny][nx] == 0:
                if (ny, nx) not in parent:
                    parent[(ny, nx)] = (y, x)
                    dq.append((ny, nx))
    # reconstruct path
    if end not in parent:
        return []
    path: list[Position] = []
    cur: Position | None = end
    while cur is not None:
        path.append(cur)
        cur = parent[cur]
    path.reverse()
    return path


def count_cheats(
    path: list[Position], min_savings: int, max_distance: int | None = None
) -> int:
    n = len(path)
    if n == 0 or min_savings >= n:
        return 0
    idx_map: dict[Position, int] = {pos: i for i, pos in enumerate(path)}

    if max_distance is None:
        offsets = [(2, 0), (-2, 0), (0, 2), (0, -2), (1, 1), (1, -1), (-1, 1), (-1, -1)]
    else:
        offsets = []
        md = max_distance
        for dy in range(-md, md + 1):
            for dx in range(-md, md + 1):
                d = abs(dx) + abs(dy)
                if 2 <= d <= md:
                    offsets.append((dy, dx))
    count = 0
    for i, (y, x) in enumerate(path):
        for dy, dx in offsets:
            coord = (y + dy, x + dx)
            j = idx_map.get(coord)
            if j is None or j <= i:
                continue
            dst = abs(dy) + abs(dx)
            if j - i - dst >= min_savings:
                count += 1
    return count
