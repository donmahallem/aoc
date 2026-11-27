import typing


def calculate_fish_counts(
    initial_state: list[int, int, int, int, int, int, int, int, int], depth: int
) -> int:
    if depth == 0:
        return sum(initial_state)
    for i in range(depth):
        current_idx = i % 9
        new_spawns = initial_state[current_idx]
        initial_state[(current_idx + 7) % 9] += new_spawns
        initial_state[current_idx] += new_spawns
        initial_state[current_idx] = 0
    return sum(initial_state)


def parse_input(
    input_data: typing.TextIO,
) -> list[int, int, int, int, int, int, int, int, int]:
    # read single line
    line = input_data.readline().strip()
    counts = [0] * 9
    for num in map(int, line.split(",")):
        counts[num] += 1
    return counts
