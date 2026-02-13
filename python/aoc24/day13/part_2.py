import typing
from .part_1 import parseInput, calc


def Part2(input: typing.TextIO) -> int:
    machines = parseInput(input)
    summe = 0
    for i, machine in enumerate(machines):
        target_x, target_y = machine["price"]
        _, btn_a_x, btn_a_y = machine["buttons"][0]
        _, btn_b_x, btn_b_y = machine["buttons"][1]
        fac_1, fac_2 = calc(
            (btn_a_x, btn_a_y),
            (btn_b_x, btn_b_y),
            (target_x + 10000000000000, target_y + 10000000000000),
        )
        if fac_1.is_integer() and fac_2.is_integer():
            summe += int(fac_1) * 3 + int(fac_2)
    return summe
