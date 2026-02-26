import typing
from .part_1 import _move_next, parseField, State, _DX, _DY


def _encode(state: State) -> int:
    y, x, d = state
    return (y << 16) | (x << 2) | d


def Part2(input: typing.TextIO) -> int:
    grid, start = parseField(input)
    rows, cols = len(grid), len(grid[0])

    obs_rows: list[list[int]] = [[] for _ in range(rows)]
    obs_cols: list[list[int]] = [[] for _ in range(cols)]
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == 1:
                obs_rows[r].append(c)
                obs_cols[c].append(r)

    base_path = []
    visited = set()
    curr = start
    while curr:
        y, x, d = curr
        if (y, x) not in visited:
            # Store the state JUST BEFORE we would have hit this tile
            base_path.append(curr)
            visited.add((y, x))
        curr = _move_next(grid, curr)

    def check_loop(obs_y, obs_x, resume_state):
        y, x, d = resume_state
        seen = set()

        while True:
            state = (y, x, d)
            if state in seen: return True
            seen.add(state)

            ny, nx = -1, -1
            if d == 0:  # Up
                walls = [r for r in obs_cols[x] if r < y]
                potential_y = max(walls) if walls else -1
                obstacle_y = obs_y if obs_x == x and obs_y < y else -1
                ny, nx = max(potential_y, obstacle_y), x
            elif d == 1:  # Right
                walls = [c for c in obs_rows[y] if c > x]
                potential_x = min(walls) if walls else cols
                obstacle_x = obs_x if obs_y == y and obs_x > x else cols
                ny, nx = y, min(potential_x, obstacle_x)
            elif d == 2:  # Down
                walls = [r for r in obs_cols[x] if r > y]
                potential_y = min(walls) if walls else rows
                obstacle_y = obs_y if obs_x == x and obs_y > y else rows
                ny, nx = min(potential_y, obstacle_y), x
            elif d == 3:  # Left
                walls = [c for c in obs_rows[y] if c < x]
                potential_x = max(walls) if walls else -1
                obstacle_x = obs_x if obs_y == y and obs_x < x else -1
                ny, nx = y, max(potential_x, obstacle_x)

            # Check if we go off-map
            if ny < 0 or ny >= rows or nx < 0 or nx >= cols:
                return False

            # Move to the cell right before the obstacle and turn
            y, x = ny - _DY[d], nx - _DX[d]
            d = (d + 1) % 4

    count = 0
    for i in range(1, len(base_path)):
        obs_y, obs_x, _ = base_path[i]
        # Resume from the state at index i-1
        if check_loop(obs_y, obs_x, base_path[i - 1]):
            count += 1

    return count
