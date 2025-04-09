package day17

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/aoc_utils"
)

func SearchForNeutralElement(reg *Register, program *Program, targetProgram *Program) bool {
	output := make([]int, 0)
	pointer := 0
	targetProgramLength := len(*targetProgram)
	var opcode, operand, operand_value, programLength int
	for pointer < len(*program) {
		programLength = len(output)
		if programLength > 0 && !slices.Equal(output, (*targetProgram)[0:aoc_utils.Min(programLength, targetProgramLength)]) {
			return false
		}
		opcode = (*program)[pointer]
		operand = (*program)[pointer+1]
		if operand < 4 {
			operand_value = operand
		} else if operand >= 4 && operand < 7 {
			operand_value = (*reg)[operand-4]
		}
		switch opcode {
		case 0:
			(*reg)[0] = (*reg)[0] / aoc_utils.IntPow(2, operand_value)
		case 1:
			(*reg)[1] = (*reg)[1] ^ operand
		case 2:
			(*reg)[1] = operand_value % 8
		case 3:
			if (*reg)[0] != 0 {
				pointer = operand_value
				continue
			}
		case 4:
			(*reg)[1] = (*reg)[1] ^ (*reg)[2]
		case 5:
			output = append(output, operand_value%8)
		case 6:
			(*reg)[1] = (*reg)[0] / (aoc_utils.IntPow(2, operand_value))
		case 7:
			(*reg)[2] = (*reg)[0] / (aoc_utils.IntPow(2, operand_value))
		}
		pointer += 2
	}
	return slices.Equal(*targetProgram, output)
}

func A(program *Program) int {
	var testRegisters Register = Register{}
	registerValue := 0
	var tempTargetProgram Program
	for i := len(*program) - 1; i >= 0; i-- {
		for {
			tempTargetProgram = append(tempTargetProgram[:0], (*program)[i:]...)
			testRegisters[0] = registerValue
			testRegisters[1] = 0
			testRegisters[2] = 0
			if SearchForNeutralElement(&testRegisters, program, &tempTargetProgram) {
				if i == 0 {
					return registerValue
				}
				registerValue *= 8
				break
			}
			registerValue++
		}
	}
	return 0
}

func Part2(in io.Reader) int {
	_, program := ParseInput(in)
	return A(&program)
}
