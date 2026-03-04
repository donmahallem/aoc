import typing
import re


def _parse_input(input: typing.TextIO) -> tuple[list[int], dict[str, list[str]]]:
    """Return (instructions, nodes) where instructions is list of 0(L) or 1(R),
    and nodes maps node_id -> [left, right]."""
    lines = [line.strip() for line in input if line.strip()]
    # first line is L/R instructions
    instructions: list[int] = [0 if ch == "L" else 1 for ch in lines[0]]
    nodes: dict[str, list[str]] = {}
    for line in lines[1:]:
        m = re.match(r"(\w+)\s*=\s*\((\w+),\s*(\w+)\)", line)
        if m:
            nodes[m.group(1)] = [m.group(2), m.group(3)]
    return instructions, nodes
