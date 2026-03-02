import typing


def parse_line(line: str) -> tuple[int, list[int], list[int]]:
    """Return (card_id, winning_numbers, picked_numbers)."""
    card_part, numbers_part = line.split(":", 1)
    card_id = int(card_part.split()[1])
    winning_part, picked_part = numbers_part.split("|", 1)
    winning = list(map(int, winning_part.split()))
    picked = list(map(int, picked_part.split()))
    return card_id, winning, picked


def count_wins(winning: list[int], picked: list[int]) -> int:
    """Count how many picked numbers are in the winning set."""
    winning_set = set(winning)
    return sum(1 for p in picked if p in winning_set)
