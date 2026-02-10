package day17

import (
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func SearchForNeutralElement(reg *Register, program *Program, targetProgram *Program) (bool, error) {
	output := make([]int, 0)
	pointer := 0
	targetProgramLength := len(*targetProgram)
	var opcode, operand, operand_value, programLength int
	if len(*program)%2 != 0 {
		return false, aoc_utils.NewParseError("malformed program: odd length", nil)
	}
	for pointer < len(*program) {
		if pointer+1 >= len(*program) {
			return false, aoc_utils.NewParseError("malformed program: missing operand", nil)
		}
		programLength = len(output)
		if programLength > 0 && !slices.Equal(output, (*targetProgram)[0:min(programLength, targetProgramLength)]) {
			return false, nil
		}
		opcode = (*program)[pointer]
		operand = (*program)[pointer+1]
		if operand < 0 {
			return false, aoc_utils.NewParseError("invalid operand", nil)
		}
		if operand < 4 {
			operand_value = operand
		} else if operand >= 4 && operand < 7 {
			operand_value = (*reg)[operand-4]
		} else {
			return false, aoc_utils.NewParseError("invalid operand", nil)
		}
		switch opcode {
		case 0:
			(*reg)[0] = (*reg)[0] / int_util.IntPow(2, operand_value)
		case 1:
			(*reg)[1] = (*reg)[1] ^ operand
		case 2:
			(*reg)[1] = operand_value % 8
		case 3:
			if (*reg)[0] != 0 {
				if operand_value < 0 || operand_value >= len(*program) {
					return false, aoc_utils.NewParseError("invalid jump target", nil)
				}
				pointer = operand_value
				continue
			}
		case 4:
			(*reg)[1] = (*reg)[1] ^ (*reg)[2]
		case 5:
			output = append(output, operand_value%8)
		case 6:
			(*reg)[1] = (*reg)[0] / (int_util.IntPow(2, operand_value))
		case 7:
			(*reg)[2] = (*reg)[0] / (int_util.IntPow(2, operand_value))
		default:
			return false, aoc_utils.NewParseError("invalid opcode", nil)
		}
		pointer += 2
	}
	return slices.Equal(*targetProgram, output), nil
}

func A(program *Program) (int, error) {
	var testRegisters Register = Register{}
	registerValue := 0
	var tempTargetProgram Program
	for i := len(*program) - 1; i >= 0; i-- {
		for {
			tempTargetProgram = append(tempTargetProgram[:0], (*program)[i:]...)
			testRegisters[0] = registerValue
			testRegisters[1] = 0
			testRegisters[2] = 0
			found, err := SearchForNeutralElement(&testRegisters, program, &tempTargetProgram)
			if err != nil {
				return 0, err
			}
			if found {
				if i == 0 {
					return registerValue, nil
				}
				registerValue *= 8
				break
			}
			registerValue++
		}
	}
	return 0, nil
}

func Part2(in io.Reader) (int, error) {
	data, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	res, err := A(&data.Program)
	if err != nil {
		return 0, err
	}
	return res, nil
}
