import typing
import heapq
from .shared import _parse_input


def Part2(input: typing.TextIO) -> int:
    data = _parse_input(input)

    # files maps file_id -> (position, size)
    files: dict[int, tuple[int, int]] = {}
    # spaces[s] is a min-heap of starting positions of free spaces of size s
    spaces: list[list[int]] = [[] for _ in range(10)]

    pos = 0
    for idx, size in enumerate(data):
        if size > 0:
            if idx % 2 == 0:
                file_id = idx // 2
                files[file_id] = (pos, size)
            else:
                spaces[size].append(pos)
        pos += size

    # initialize the min-heaps for spaces
    for s in range(10):
        heapq.heapify(spaces[s])

    max_id = len(data) // 2
    for file_id in range(max_id, -1, -1):
        if file_id not in files:
            continue

        old_pos, size = files[file_id]

        # Find the leftmost space that can fit the file
        best_size = -1
        best_pos = old_pos
        for s in range(size, 10):
            if spaces[s] and spaces[s][0] < best_pos:
                best_pos = spaces[s][0]
                best_size = s

        if best_size != -1:
            new_pos = heapq.heappop(spaces[best_size])
            files[file_id] = (new_pos, size)

            # The remainder of the space is added back to the pool
            if best_size > size:
                heapq.heappush(spaces[best_size - size], new_pos + size)

    total = 0
    for file_id, (final_pos, size) in files.items():
        total += file_id * (final_pos * size + (size * (size - 1)) // 2)

    return total
