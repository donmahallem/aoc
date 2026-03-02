import typing
from dataclasses import dataclass


@dataclass
class Block:
    red: int = 0
    green: int = 0
    blue: int = 0


def parse_line(line: str) -> tuple[int, list[Block]]:
    """Parse 'Game N: ...;...' into game id and list of Block draws."""
    game_part, draws_part = line.split(":", 1)
    game_id = int(game_part.split()[1])
    blocks: list[Block] = []
    for draw in draws_part.split(";"):
        block = Block()
        for part in draw.split(","):
            part = part.strip()
            count_str, colour = part.split()
            count = int(count_str)
            if colour == "red":
                block.red = count
            elif colour == "green":
                block.green = count
            elif colour == "blue":
                block.blue = count
        blocks.append(block)
    return game_id, blocks
