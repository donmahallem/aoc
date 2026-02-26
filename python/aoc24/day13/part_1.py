import typing
import re

Button: typing.TypeAlias = tuple[str, int, int]


class Machine(typing.TypedDict):
    buttons: list[Button]
    price: tuple[int, int]


Machines: typing.TypeAlias = list[Machine]


def parseInput(input: typing.TextIO) -> Machines:
    machines: Machines = list()
    machine: Machine = Machine(buttons=list(), price=(0, 0))
    button_regex = re.compile(
        r"(?:Button\s([A-Za-z]+)\:)\s(?:[XY]([+-]\d+)),\s(?:[XY]([+-]\d+))")
    prize_regex = re.compile(r"(?:Prize)\:\s(?:[XY]=(\d+)),\s*(?:[XY]=(\d+))")
    for line in input.readlines():
        line = line.strip()
        if len(line) == 0:
            continue
        reg_res = button_regex.match(line)
        if reg_res:
            button: Button = (
                reg_res.groups()[0],
                int(reg_res.groups()[1]),
                int(reg_res.groups()[2]),
            )
            machine["buttons"].append(button)
        reg_res = prize_regex.match(line)
        if reg_res:
            machine["price"] = (int(reg_res.groups()[0]),
                                int(reg_res.groups()[1]))
            machines.append(machine)
            machine = Machine(buttons=list(), price=(0, 0))
    return machines


Vector = tuple[float, float]


def calc(v1: Vector, v2: Vector, target) -> Vector:
    x_1, y_1 = v1
    x_2, y_2 = v2
    x_3, y_3 = target
    v2_factor = ((x_3 * y_1) - (x_1 * y_3)) / ((x_2 * y_1) - (x_1 * y_2))
    v1_factor = (x_3 - (v2_factor * x_2)) / x_1
    return (v1_factor, v2_factor)


def Part1(input: typing.TextIO) -> int:
    machines = parseInput(input)
    summe: int = 0
    for i, machine in enumerate(machines):
        target_x, target_y = machine["price"]
        _, btn_a_x, btn_a_y = machine["buttons"][0]
        _, btn_b_x, btn_b_y = machine["buttons"][1]
        fac_1, fac_2 = calc((btn_a_x, btn_a_y), (btn_b_x, btn_b_y),
                            (target_x, target_y))
        if fac_1.is_integer() and fac_2.is_integer():
            summe += int(fac_1) * 3 + int(fac_2)
    return int(summe)
