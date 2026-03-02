import typing
from dataclasses import dataclass


@dataclass(slots=True)
class Block:
    red: int = 0
    green: int = 0
    blue: int = 0


def _parse_line(line: str) -> tuple[int, list[Block]]:
    game_part, draws_part = line.split(":", 1)

    # 2. Direct string slicing instead of splitting again
    game_id = int(game_part[5:])

    blocks: list[Block] = []
    for draw in draws_part.split(";"):
        block = Block()
        for part in draw.split(","):
            # 3. .split() ignores whitespace by default, making .strip() redundant
            count_str, colour = part.split()
            count = int(count_str)

            # 4. Micro-optimization: Check only the first character
            if colour[0] == "r":
                block.red = count
            elif colour[0] == "g":
                block.green = count
            else:
                block.blue = count

        blocks.append(block)

    return game_id, blocks
