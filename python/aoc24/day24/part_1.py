import typing
from .operation import Operation


def parse_input(input: typing.TextIO)->tuple[list,list[Operation]] :
    data="\n".join([ln.strip() for ln in input.readlines()])
    register=list()
    wires=list()
    register, wires = data.split("\n\n")

    register = {
        line.strip().split(":")[0]: int(line.strip().split(":")[1])
        for line in register.split("\n")
    }
    wires = [line.strip() for line in wires.split("\n")]
    wires = [
        Operation([item[0],item[2]],item[1],item[4]) for item in [line.split(" ") for line in wires]
    ]
    return register,wires

def combine(registers,wires:list[Operation]):
    connected = True
    while connected:
        connected = False
        for line in wires:
            if line.output in registers:
                continue
            if line.inputs[0] in registers and line.inputs[1] in registers:
                val_a = registers[line.inputs[0]]
                val_b = registers[line.inputs[1]]
                if line.operator == "AND":
                    out = val_a and val_b
                elif line.operator == "OR":
                    out = val_a or val_b
                elif line.operator == "XOR":
                    out = val_a ^ val_b
                registers[line.output] = out
                connected = True

    sorted_keys = sorted(registers.keys())

    output_value = 0
    for key in sorted_keys:
        if key[0] == "z":
            if registers[key] > 0:
                key_val = int(key[1:])
                output_value += 1 << key_val
    return output_value

def Part1(input: typing.TextIO) -> int:
    registers,wires = parse_input(input)
    return combine(registers,wires)
