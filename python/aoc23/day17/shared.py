import typing
import heapq
from util.point import Point

type _CellMap = dict[Point, int]


def _parse_input(input: typing.TextIO) -> tuple[_CellMap, int, int]:
    """Return a dictionary mapping Point coordinates to digit values."""
    cells: _CellMap = {}
    height = 0
    expected_width: int = -1
    for line in input:
        line = line.strip()
        if not line:
            continue
        if expected_width < 0:
            expected_width = len(line)
        elif len(line) != expected_width:
            raise ValueError(f"Expected width {expected_width} but got {len(line)}")
        for ch, x in zip(line, range(expected_width)):
            cells[Point(x, height)] = int(ch)
        height += 1
    return cells, expected_width, height


_DIRS = [Point(0, 1), Point(1, 0), Point(0, -1), Point(-1, 0)]


def _find_shortest_path(
    cells: _CellMap, width: int, height: int, min_straight: int, max_straight: int
) -> int:
    """Dijkstra with state (cost, dir_idx, steps, pos)."""
    if width == 0 or height == 0:
        return 0
    target = Point(width - 1, height - 1)

    # state -> best cost
    best: dict[tuple[Point, int, int], int] = {}

    # state: (cost, dir_idx, steps, pos)
    # Point supports __lt__ so heapq comparisons are safe.
    heap: list[tuple[int, int, int, Point]] = []
    # Start from (0,0) - push initial moves in all 4 directions
    for dir_idx, dir in enumerate(_DIRS):
        next_pos = Point(0, 0) + dir
        if 0 <= next_pos.x < width and 0 <= next_pos.y < height:
            cost = cells[next_pos]
            state = (next_pos, dir_idx, 1)
            if state not in best or cost < best[state]:
                best[state] = cost
                heapq.heappush(heap, (cost, dir_idx, 1, next_pos))

    while heap:
        cost, dir_idx, steps, pos = heapq.heappop(heap)

        state = (pos, dir_idx, steps)
        if best.get(state, cost + 1) < cost:
            continue

        if pos == target:
            if steps >= min_straight:
                return cost
            continue

        if steps < min_straight:
            # must continue straight
            dir = _DIRS[dir_idx]
            next_pos = pos + dir
            if 0 <= next_pos.x < width and 0 <= next_pos.y < height:
                new_cost = cost + cells[next_pos]
                new_state = (next_pos, dir_idx, steps + 1)
                if new_cost < best.get(new_state, new_cost + 1):
                    best[new_state] = new_cost
                    heapq.heappush(heap, (new_cost, dir_idx, steps + 1, next_pos))
        else:
            # can turn or continue (if under max_straight)
            for new_dir in range(4):
                # can't reverse
                if (new_dir + 2) % 4 == dir_idx:
                    continue
                new_steps = steps + 1 if new_dir == dir_idx else 1
                if new_dir == dir_idx and steps >= max_straight:
                    continue
                dir = _DIRS[new_dir]
                next_pos = pos + dir
                if 0 <= next_pos.x < width and 0 <= next_pos.y < height:
                    new_cost = cost + cells[next_pos]
                    new_state = (next_pos, new_dir, new_steps)
                    if new_cost < best.get(new_state, new_cost + 1):
                        best[new_state] = new_cost
                        heapq.heappush(heap, (new_cost, new_dir, new_steps, next_pos))

    return 0
