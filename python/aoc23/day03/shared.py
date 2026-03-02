import typing
from dataclasses import dataclass, field


@dataclass
class Number:
    value: int
    start_x: int
    end_x: int  # inclusive
    y: int


@dataclass
class Part:
    char: str
    x: int
    y: int


def find_objects(grid: list[str]) -> tuple[list[Part], list[Number]]:
    """Scan the grid and return all parts (symbols) and numbers."""
    parts: list[Part] = []
    numbers: list[Number] = []

    for y, row in enumerate(grid):
        x = 0
        while x < len(row):
            ch = row[x]
            if ch.isdigit():
                start_x = x
                num_str = ""
                while x < len(row) and row[x].isdigit():
                    num_str += row[x]
                    x += 1
                numbers.append(Number(int(num_str), start_x, x - 1, y))
            elif ch != ".":
                parts.append(Part(ch, x, y))
                x += 1
            else:
                x += 1
    return parts, numbers


def find_adjacent_numbers(p: Part, numbers: list[Number]) -> list[Number]:
    """Return all numbers that are adjacent (including diagonally) to the given part."""
    adjacent: list[Number] = []
    for num in numbers:
        if abs(num.y - p.y) <= 1 and num.start_x - 1 <= p.x <= num.end_x + 1:
            adjacent.append(num)
    return adjacent
