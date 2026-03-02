import sys
import typing

_CARD_RANKS_P1: dict[str, int] = {
    "2": 0,
    "3": 1,
    "4": 2,
    "5": 3,
    "6": 4,
    "7": 5,
    "8": 6,
    "9": 7,
    "T": 8,
    "J": 9,
    "Q": 10,
    "K": 11,
    "A": 12,
}


def _rating_p1(freq_sorted: list[int]) -> int:
    top = freq_sorted[-1]
    if top == 5:
        return 7
    if top == 4:
        return 6
    if top == 3:
        return 5 if freq_sorted[-2] == 2 else 4
    if top == 2:
        return 3 if freq_sorted[-2] == 2 else 2
    return 1


def _hand_hash(rating: int, hand_ranks: list[int]) -> int:
    h = rating
    for r in hand_ranks:
        h = h * 13 + r
    return h


def _parse_input_p1(input: typing.TextIO) -> list[tuple[int, int]]:
    games: list[tuple[int, int]] = []
    for line in input:
        line = line.strip()
        if not line:
            continue
        hand_str, bid_str = line.split()
        freq: dict[int, int] = {}
        hand_ranks: list[int] = []
        for ch in hand_str:
            r = _CARD_RANKS_P1[ch]
            hand_ranks.append(r)
            freq[r] = freq.get(r, 0) + 1
        freq_sorted = sorted(freq.values())
        rating = _rating_p1(freq_sorted)
        games.append((_hand_hash(rating, hand_ranks), int(bid_str)))
    return games


def Part1(input: typing.TextIO) -> int:
    games = _parse_input_p1(input)
    games.sort(key=lambda g: g[0])
    return sum(bid * (idx + 1) for idx, (_, bid) in enumerate(games))


if __name__ == "__main__":
    print(Part1(sys.stdin))
