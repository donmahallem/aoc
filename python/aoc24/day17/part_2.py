import typing
from .part_1 import parseField


def checkNum2(base_program, target, target_program):
    test_register = [target, 0, 0]
    pointer = 0
    output = []
    while pointer < len(base_program):
        if output != target_program[0 : len(output)]:
            # print(output,program)
            return False
        opcode = base_program[pointer]
        operand = base_program[pointer + 1]
        if operand < 4:
            operand_value = operand
        elif operand >= 4 and operand < 7:
            operand_value = test_register[operand - 4]
        ######
        if opcode == 0:
            test_register[0] = int(test_register[0] / (2**operand_value))
        elif opcode == 1:
            test_register[1] = test_register[1] ^ operand
        elif opcode == 2:
            test_register[1] = operand_value % 8
        elif opcode == 3:
            if test_register[0] != 0:
                pointer = operand_value
                continue
        elif opcode == 4:
            test_register[1] = test_register[1] ^ test_register[2]
        elif opcode == 5:
            output.append(operand_value % 8)
        elif opcode == 6:
            test_register[1] = int(test_register[0] / (2**operand_value))
        elif opcode == 7:
            test_register[2] = int(test_register[0] / (2**operand_value))
        pointer += 2

    return output


def Part2(input: typing.TextIO) -> int:
    register, program = parseField(input)
    steps = 1
    j = 0
    for i in range(len(program) - 1, -1, -1):
        while True:
            result = checkNum2(program, j, program[i : len(program)])
            if result == program[i : len(program)]:
                if i == 0:
                    return j
                j *= 8
                break
            j += 1
    return -1
