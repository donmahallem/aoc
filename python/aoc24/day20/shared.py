import typing
import numpy as np
from collections import deque
from dataclasses import dataclass
from util.typed_coord import Coord, CoordX, CoordY


@dataclass(slots=True)
class _Field:
    grid: list[str]
    width: CoordX
    height: CoordY
    start: Coord
    end: Coord


def _parse_field(input: typing.TextIO) -> _Field | None:
    grid = [line.rstrip("\n") for line in input]
    if not grid:
        return None

    h: CoordY = len(grid)
    w: CoordX = len(grid[0])

    start: Coord | None = None
    end: Coord | None = None

    for y, line in enumerate(grid):
        if start is None:
            x = line.find("S")
            if x != -1:
                start = (y, x)
        if end is None:
            x = line.find("E")
            if x != -1:
                end = (y, x)

    if start is None or end is None:
        return None

    return _Field(grid=grid, width=w, height=h, start=start, end=end)


def _compute_shortest_path(field: _Field) -> list[Coord]:
    w, h = field.width, field.height
    grid = field.grid
    sy, sx = field.start
    ey, ex = field.end

    start_idx = sy * w + sx
    end_idx = ey * w + ex

    parent = [-1] * (w * h)
    parent[start_idx] = start_idx

    dq = deque([start_idx])
    dirs = (-w, w, -1, 1)

    while dq:
        idx = dq.popleft()
        if idx == end_idx:
            break

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
        path.append((py, px))
        cur_idx = parent[cur_idx]

    path.append(field.start)
    path.reverse()
    return path


def _count_cheats(
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

    path_y = np.array([p[0] for p in path], dtype=np.int32)
    path_x = np.array([p[1] for p in path], dtype=np.int32)

    pad = md
    grid_w = int(path_x.max()) + 1 + 2 * pad
    grid_h = int(path_y.max()) + 1 + 2 * pad

    idx_map = np.full(grid_w * grid_h, -1, dtype=np.int32)
    flat_indices = (path_y + pad) * grid_w + (path_x + pad)
    idx_map[flat_indices] = np.arange(n, dtype=np.int32)

    n_check = n - min_savings
    centers = flat_indices[:n_check, np.newaxis]
    i_vals = np.arange(n_check, dtype=np.int32)[:, np.newaxis]

    deltas = np.array(all_dy, dtype=np.int32) * grid_w + np.array(
        all_dx, dtype=np.int32
    )
    distances = np.array(all_dist, dtype=np.int32)

    all_j = idx_map[centers + deltas]

    thresholds = i_vals + distances + min_savings

    return int(np.count_nonzero(all_j >= thresholds))
