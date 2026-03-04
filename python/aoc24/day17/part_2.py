import typing
from .part_1 import parseField


def checkNum2(base_program, target, target_program):
    register_a, register_b, register_c = target, 0, 0
    pointer = 0
    out_idx = 0
    target_len = len(target_program)

    prog_len = len(base_program)

    while pointer < prog_len:
        opcode = base_program[pointer]
        operand = base_program[pointer + 1]

        if opcode in {0, 2, 5, 6, 7}:
            if operand <= 3:
                combo = operand
            elif operand == 4:
                combo = register_a
            elif operand == 5:
                combo = register_b
            elif operand == 6:
                combo = register_c

        match opcode:
            case 0:
                register_a >>= combo
            case 1:
                register_b ^= operand
            case 2:
                register_b = combo % 8
            case 3:
                if register_a != 0:
                    pointer = operand
                    continue
            case 4:
                register_b ^= register_c
            case 5:
                val = combo % 8
                if out_idx >= target_len or val != target_program[out_idx]:
                    return False
                out_idx += 1
            case 6:
                register_b = register_a >> combo
            case 7:
                register_c = register_a >> combo

        pointer += 2

    return out_idx == target_len


def Part2(input_stream: typing.TextIO) -> int:
    _, program = parseField(input_stream)
    j = 0
    program_len = len(program)

    for i in range(program_len - 1, -1, -1):
        target_suffix = program[i:]
        while True:
            if checkNum2(program, j, target_suffix):
                if i == 0:
                    return j
                j <<= 3
                break
            j += 1
    return -1
