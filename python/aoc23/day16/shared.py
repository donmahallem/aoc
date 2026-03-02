import typing

_DIR_RIGHT = complex(1, 0)
_DIR_LEFT = complex(-1, 0)
_DIR_UP = complex(0, -1)
_DIR_DOWN = complex(0, 1)

_DIR_BIT: dict[complex, int] = {
    _DIR_RIGHT: 1,
    _DIR_LEFT: 2,
    _DIR_UP: 4,
    _DIR_DOWN: 8,
}

type _Field = dict[complex, str]


def parse_input(input: typing.TextIO) -> tuple[_Field, int, int]:
    expected_width: int = -1
    height: int = 0
    field: _Field = {}
    for line in input:
        line = line.strip()
        if not line:
            continue
        if expected_width < 0:
            expected_width = len(line)
        elif len(line) != expected_width:
            raise ValueError(
                f"Inconsistent width at line {height}: expected {expected_width}, got {len(line)}"
            )
        for x, ch in enumerate(line):
            field[complex(x, height)] = ch
        height += 1
    return field, expected_width, height


def simulate(
    field: _Field,
    width: int,
    height: int,
    start_pos: complex,
    start_dir: complex,
) -> int:
    """Simulate beam and return number of energized cells."""
    memory: dict[complex, int] = {}
    stack: list[tuple[complex, complex]] = [(start_pos, start_dir)]

    while stack:
        pos, d = stack.pop()

        while True:
            if not (0 <= pos.real < width and 0 <= pos.imag < height):
                break
            bit = _DIR_BIT[d]
            if memory.get(pos, 0) & bit:
                break
            memory[pos] = memory.get(pos, 0) | bit

            cell = field[pos]

            if cell == "/":
                # reflect: (dx,dy) -> (-dy,-dx)
                d = complex(-d.imag, -d.real)
            elif cell == "\\":
                # reflect: (dx,dy) -> (dy,dx)
                d = complex(d.imag, d.real)
            elif cell == "|":
                if d.real != 0:  # moving horizontally
                    stack.append((pos + _DIR_UP, _DIR_UP))
                    d = _DIR_DOWN
            elif cell == "-":
                if d.imag != 0:  # moving vertically
                    stack.append((pos + _DIR_LEFT, _DIR_LEFT))
                    d = _DIR_RIGHT

            pos += d

    return len(memory)
