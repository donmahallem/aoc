import typing


def convert_to_bit_map(inp: str) -> int:
    result = 0
    for char in inp:
        result |= 1 << (ord(char) - ord("a"))
    return result


class SourceData(typing.NamedTuple):
    mapping: list[int]
    output: list[int]


def parse_input(
    input_data: typing.TextIO,
) -> list[SourceData]:
    result = []
    for line in input_data:
        line = line.strip()
        if not line:
            continue
        splitted = line.split(" | ")
        assert splitted and len(splitted) == 2

        mapping = [
            convert_to_bit_map(x)
            for x in sorted([x.strip() for x in splitted[0].split()], key=len)
        ]
        output = [convert_to_bit_map(x.strip()) for x in splitted[1].split()]
        result.append(SourceData(mapping=mapping, output=output))
    return result


def decode_mapping(mapping: list[int]) -> dict[int, int]:

    one = mapping[0]
    seven = mapping[1]
    four = mapping[2]
    all_5 = mapping[3] & mapping[4] & mapping[5]
    all_6 = mapping[6] & mapping[7] & mapping[8]
    eight = mapping[9]

    seg_a = seven - one
    seg_f = one & all_6
    seg_c = one - seg_f
    seg_d = (four - one) & all_5
    seg_b = (four - one) - seg_d
    seg_g = all_5 - seg_a - seg_d
    seg_e = eight - four - seg_a - seg_g

    return {
        seg_a | seg_b | seg_c | seg_e | seg_f | seg_g: 0,
        seg_c | seg_f: 1,
        seg_a | seg_c | seg_d | seg_e | seg_g: 2,
        seg_a | seg_c | seg_d | seg_f | seg_g: 3,
        seg_b | seg_c | seg_d | seg_f: 4,
        seg_a | seg_b | seg_d | seg_f | seg_g: 5,
        seg_a | seg_b | seg_d | seg_e | seg_f | seg_g: 6,
        seg_a | seg_c | seg_f: 7,
        seg_a | seg_b | seg_c | seg_d | seg_e | seg_f | seg_g: 8,
        seg_a | seg_b | seg_c | seg_d | seg_f | seg_g: 9,
    }
