import typing
import numpy as np
from collections import deque
from dataclasses import dataclass
from util.typed_coord import Coord, CoordX, CoordY


@dataclass(slots=True)
class _Field:
    grid: list[str]
    width: int
    height: int
    start: Coord
    end: Coord


def parseField(input: typing.TextIO) -> _Field | None:
    grid = [line.rstrip("\n") for line in input]
    if not grid:
        return None

    h = len(grid)
    w = len(grid[0])

    start: Coord | None = None
    end: Coord | None = None

    for y, line in enumerate(grid):
        if start is None:
            x = line.find("S")
            if x != -1:
                start = Coord((CoordY(y), CoordX(x)))
        if end is None:
            x = line.find("E")
            if x != -1:
                end = Coord((CoordY(y), CoordX(x)))

    if start is None or end is None:
        return None

    return _Field(grid=grid, width=w, height=h, start=start, end=end)


def compute_shortest_path(field: _Field) -> list[Coord]:
    w, h = field.width, field.height
    grid = field.grid
    sy, sx = field.start
    ey, ex = field.end

    start_idx = sy * w + sx
    end_idx = ey * w + ex

    # Track the parent using flat 1D indices (-1 means unvisited)
    parent = [-1] * (w * h)
    parent[start_idx] = start_idx

    dq = deque([start_idx])
    dirs = (-w, w, -1, 1)

    while dq:
        idx = dq.popleft()
        if idx == end_idx:
            break

        y, x = divmod(idx, w)

        for delta in dirs:
            ny, nx = divmod(idx + delta, w)
            # Bound checks natively derived from our 2D boundaries
            if 0 <= ny < h and 0 <= nx < w:
                nidx = idx + delta
                if grid[ny][nx] != "#" and parent[nidx] == -1:
                    parent[nidx] = idx
                    dq.append(nidx)

    if parent[end_idx] == -1:
        return []

    path: list[Coord] = []
    cur_idx = end_idx
    while cur_idx != start_idx:
        py, px = divmod(cur_idx, w)
        path.append(Coord((CoordY(py), CoordX(px))))
        cur_idx = parent[cur_idx]

    path.append(field.start)
    path.reverse()
    return path


def count_cheats(
    path: list[Coord], min_savings: int, max_distance: int | None = None
) -> int:
    n = len(path)
    if n == 0 or min_savings >= n:
        return 0

    md = 2 if max_distance is None else max_distance

    # Create a flat array of all valid offsets and their manhattan distances
    all_dy, all_dx, all_dist = [], [], []
    for dy in range(-md, md + 1):
        for dx in range(-md, md + 1):
            d = abs(dx) + abs(dy)
            if 2 <= d <= md:
                all_dy.append(dy)
                all_dx.append(dx)
                all_dist.append(d)

    # Convert path to numpy arrays
    path_y = np.array([p[0] for p in path], dtype=np.int32)
    path_x = np.array([p[1] for p in path], dtype=np.int32)

    # Pad the grid by max_distance to eliminate bounds checking
    pad = md
    grid_w = int(path_x.max()) + 1 + 2 * pad
    grid_h = int(path_y.max()) + 1 + 2 * pad

    # 1D numpy array mapping padded coordinates to path indices (-1 = not on path)
    idx_map = np.full(grid_w * grid_h, -1, dtype=np.int32)
    flat_indices = (path_y + pad) * grid_w + (path_x + pad)
    idx_map[flat_indices] = np.arange(n, dtype=np.int32)

    # Only check positions that can possibly yield valid cheats
    n_check = n - min_savings
    centers = flat_indices[:n_check, np.newaxis]
    i_vals = np.arange(n_check, dtype=np.int32)[:, np.newaxis]

    # Convert offsets to flat deltas
    deltas = np.array(all_dy, dtype=np.int32) * grid_w + np.array(
        all_dx, dtype=np.int32
    )
    distances = np.array(all_dist, dtype=np.int32)

    # ONE giant vectorized operation:
    # all_j is shape (n_check, num_offsets) containing the path index of the destination
    all_j = idx_map[centers + deltas]

    # Valid cheat if j >= i + dist + min_savings
    thresholds = i_vals + distances + min_savings

    return int(np.count_nonzero(all_j >= thresholds))
