import typing
import itertools
import codecs


def parseInput(input: typing.TextIO) -> dict[str, list[str]]:
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
        tuple(item[0:3] + [item[4]]) for item in [line.split(" ") for line in wires]
    ]
    return register,wires


def combine(registers,wires):
    connected = True
    while connected:
        connected = False
        for line in wires:
            if line[3] in registers:
                continue
            if line[0] in registers and line[2] in registers:
                val_a = registers[line[0]]
                val_b = registers[line[2]]
                if line[1] == "AND":
                    out = val_a and val_b
                elif line[1] == "OR":
                    out = val_a or val_b
                elif line[1] == "XOR":
                    out = val_a ^ val_b
                registers[line[3]] = out
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
    registers,wires = parseInput(input)
    return combine(registers,wires)
