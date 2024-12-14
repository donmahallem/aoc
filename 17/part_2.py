import codecs

test_data = False
with codecs.open("data.txt" if test_data else "data2.txt", encoding="utf8") as f:
    data = [line.strip() for line in f.readlines()]

register = [0, 0, 0]
program = []
for line in data:
    if line.startswith("Register"):
        line_data = line.split(" ")
        register[ord(line_data[1][0]) - ord("A")] = int(line_data[2])
    if line.startswith("Program:"):
        program = [int(part) for part in line.split(" ")[1].split(",")]
print(register, program)


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


steps = 1
j = 0
for i in range(len(program) - 1, -1, -1):
    while True:
        result = checkNum2(program, j, program[i : len(program)])
        if result == program[i : len(program)]:
            print("Result:" if i == 0 else "Step:", j, result)
            j *= 8
            break
        j += 1
