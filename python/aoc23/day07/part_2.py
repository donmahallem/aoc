import sys
import typing

_CARD_RANKS_P2: dict[str, int] = {
    "J": 0,
    "2": 1,
    "3": 2,
    "4": 3,
    "5": 4,
    "6": 5,
    "7": 6,
    "8": 7,
    "9": 8,
    "T": 9,
    "Q": 10,
    "K": 11,
    "A": 12,
}


def _rating_p2(freq_sorted: list[int], jokers: int) -> int:
    top = freq_sorted[-1] + jokers if freq_sorted else jokers
    second = freq_sorted[-2] if len(freq_sorted) >= 2 else 0

    if top == 5:
        return 7
    if top == 4:
        return 6
    if top == 3:
        return 5 if second == 2 else 4
    if top == 2:
        return 3 if second == 2 else 2
    return 1


def _hand_hash(rating: int, hand_ranks: list[int]) -> int:
    h = rating
    for r in hand_ranks:
        h = h * 13 + r
    return h


def _parse_input_p2(input: typing.TextIO) -> list[tuple[int, int]]:
    games: list[tuple[int, int]] = []
    for line in input:
        line = line.strip()
        if not line:
            continue
        hand_str, bid_str = line.split()
        freq: dict[int, int] = {}
        hand_ranks: list[int] = []
        jokers = 0
        for ch in hand_str:
            r = _CARD_RANKS_P2[ch]
            hand_ranks.append(r)
            if r == 0:  # joker
                jokers += 1
            else:
                freq[r] = freq.get(r, 0) + 1
        freq_sorted = sorted(freq.values())
        rating = _rating_p2(freq_sorted, jokers)
        games.append((_hand_hash(rating, hand_ranks), int(bid_str)))
    return games


def Part2(input: typing.TextIO) -> int:
    games = _parse_input_p2(input)
    games.sort(key=lambda g: g[0])
    return sum(bid * (idx + 1) for idx, (_, bid) in enumerate(games))


if __name__ == "__main__":
    print(Part2(sys.stdin))
