import typing


def parse_input(
    input_data: typing.TextIO,
) -> list[list[int]]:
    expected_row_length = None
    rows = []
    for line in input_data:
        line = line.strip()
        if not line:
            continue
        if expected_row_length is None:
            expected_row_length = len(line)
        elif len(line) != expected_row_length:
            raise ValueError(
                f"Unexpected row length: {len(line)} (expected {expected_row_length})"
            )
        rows.append([ord(char) - ord("0") for char in line])
    return rows
