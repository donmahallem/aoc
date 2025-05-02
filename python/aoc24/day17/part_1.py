import typing


def parseField(input: typing.TextIO):
    data = list(map(str.strip, input.readlines()))

    register = [0, 0, 0]
    program = []
    for line in data:
        if line.startswith("Register"):
            line_data = line.split(" ")
            register[ord(line_data[1][0]) - ord("A")] = int(line_data[2])
        if line.startswith("Program:"):
            program = [int(part) for part in line.split(" ")[1].split(",")]
    return register, program


def Part1(input: typing.TextIO) -> list[int]:
    register, program = parseField(input)
    pointer: int = 0
    output: list[int] = []
    while pointer < len(program):
        opcode = program[pointer]
        operand = program[pointer + 1]
        if operand < 4:
            operand_value = operand
        elif operand >= 4 and operand < 7:
            operand_value = register[operand - 4]
        if opcode == 0:
            register[0] = int(register[0] / (2**operand_value))
        elif opcode == 1:
            register[1] = register[1] ^ operand
        elif opcode == 2:
            register[1] = operand_value % 8
        elif opcode == 3:
            if register[0] != 0:
                pointer = operand_value
                continue
        elif opcode == 4:
            register[1] = register[1] ^ register[2]
        elif opcode == 5:
            output.append(operand_value % 8)
        elif opcode == 6:
            register[1] = int(register[0] / (2**operand_value))
        elif opcode == 7:
            register[2] = int(register[0] / (2**operand_value))
        pointer += 2
    return output
