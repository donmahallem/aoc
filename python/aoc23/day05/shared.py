from collections import defaultdict
import typing


def _parse_input(
    input_data: typing.TextIO,
) -> tuple[list[int], list[list[tuple[int, int, int]]]]:
    seeds: list[int] | None = None
    mapData: defaultdict[str, list[tuple[int, int, int]]] = defaultdict(list)
    for line in input_data:
        line = line.strip()
        if not line:
            continue
        elif line.startswith("seeds"):
            seeds = list(map(int, line.split()[1:]))
            break
    currentGroup: list[tuple[int, int, int]] | None = None
    for line in input_data:
        line = line.strip()
        if not line:
            continue
        line_data: list[str] = line.split()
        match line_data[0]:
            case "seed-to-soil":
                currentGroup = mapData[line_data[0]]
            case "soil-to-fertilizer":
                currentGroup = mapData[line_data[0]]
            case "fertilizer-to-water":
                currentGroup = mapData[line_data[0]]
            case "water-to-light":
                currentGroup = mapData[line_data[0]]
            case "light-to-temperature":
                currentGroup = mapData[line_data[0]]
            case "temperature-to-humidity":
                currentGroup = mapData[line_data[0]]
            case "humidity-to-location":
                currentGroup = mapData[line_data[0]]
            case _:
                values = tuple(map(int, line_data))
                if len(values) != 3:
                    raise ValueError(f"Expected 3 integers in line: {line}")
                if currentGroup is None:
                    raise ValueError(f"Data line before any key: {line}")
                currentGroup.append(values)

    key_sequence = [
        "seed-to-soil",
        "soil-to-fertilizer",
        "fertilizer-to-water",
        "water-to-light",
        "light-to-temperature",
        "temperature-to-humidity",
        "humidity-to-location",
    ]
    transform_sequence: list[list[tuple[int, int, int]]] = []
    for key in key_sequence:
        if key not in mapData:
            raise ValueError(f"Missing key in input: {key}")
        transform_sequence.append(mapData[key])

    return seeds or [], transform_sequence
