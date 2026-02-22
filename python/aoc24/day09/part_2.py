import typing
from .part_1 import parseInput
import bisect


def expandToGroups(data: list[int]) -> tuple[list[tuple[int, int]], int]:
    line: list[tuple[int, int]] = []
    block_num = 0
    for idx, val in enumerate(data):
        if idx % 2 == 0:
            line.append((block_num, val))
            block_num += 1
        elif val > 0:
            line.append((-1, val))
    return line, block_num - 1


def handleRow(line: list[tuple[int, int]], current_idx: int,
              group_positions: list[int],
              blank_positions: list[int]) -> list[tuple[int, int]]:
    end_idx = group_positions[current_idx]
    if end_idx < 0:
        return line
    group_len = line[end_idx][1]
    ln = line
    chosen = -1
    for idx in blank_positions:
        if idx >= end_idx:
            break
        if ln[idx][1] >= group_len:
            chosen = idx
            break
    if chosen < 0:
        return ln
    blank_len = ln[chosen][1]
    if blank_len == group_len:
        # simple swap
        ln[chosen] = (current_idx, group_len)
        # remove this blank from list
        pos = bisect.bisect_left(blank_positions, chosen)
        if pos < len(blank_positions) and blank_positions[pos] == chosen:
            blank_positions.pop(pos)
        # convert old group to blank
        ln[end_idx] = (-1, group_len)
        bisect.insort(blank_positions, end_idx)
        group_positions[current_idx] = chosen
        return ln
    ln[chosen] = (current_idx, group_len)
    ln.insert(chosen + 1, (-1, blank_len - group_len))
    pos = bisect.bisect_left(blank_positions, chosen)
    if pos < len(blank_positions) and blank_positions[pos] == chosen:
        blank_positions[pos] = chosen + 1
    for i in range(len(blank_positions)):
        if blank_positions[i] >= chosen + 1 and i != pos:
            blank_positions[i] += 1
    # update group_positions for groups after insertion
    for j in range(len(group_positions)):
        if group_positions[j] > chosen:
            group_positions[j] += 1
    if chosen < end_idx:
        end_idx += 1
    ln[end_idx] = (-1, group_len)
    bisect.insort(blank_positions, end_idx)
    group_positions[current_idx] = chosen
    return ln


def Part2(input: typing.TextIO) -> int:
    data = parseInput(input)
    line_data, max_num = expandToGroups(data)
    group_positions = [-1] * (max_num + 1)
    blank_positions: list[int] = []
    for idx, (t, _) in enumerate(line_data):
        if t >= 0:
            group_positions[t] = idx
        else:
            blank_positions.append(idx)

    for i in range(max_num, 0, -1):
        line_data = handleRow(line_data, i, group_positions, blank_positions)
    # compute weighted sum directly
    total = 0
    idx = 0
    for val, length in line_data:
        if val >= 0:
            total += val * ((idx + idx + length - 1) * length // 2)
        idx += length
    return total
