package day17

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func CalculateOutput(reg *Register, program *Program) (*Program, error) {
	output := make(Program, 0)
	pointer := 0
	var opcode, operand, operand_value int
	if len(*program)%2 != 0 {
		return nil, aoc_utils.NewParseError("malformed program: odd length", nil)
	}
	for pointer < len(*program) {
		if pointer+1 >= len(*program) {
			return nil, aoc_utils.NewParseError("malformed program: missing operand", nil)
		}
		opcode = (*program)[pointer]
		operand = (*program)[pointer+1]
		// validate operand
		if operand < 0 {
			return nil, aoc_utils.NewParseError("invalid operand", nil)
		}
		if operand < 4 {
			operand_value = operand
		} else if operand >= 4 && operand < 7 {
			operand_value = (*reg)[operand-4]
		} else {
			return nil, aoc_utils.NewParseError("invalid operand", nil)
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
					return nil, aoc_utils.NewParseError("invalid jump target", nil)
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
			return nil, aoc_utils.NewParseError("invalid opcode", nil)
		}
		pointer += 2
	}
	return &output, nil
}

func Part1(in io.Reader) ([]int, error) {
	data, err := parseInput(in)
	if err != nil {
		return nil, err
	}
	out, err := CalculateOutput(&data.Register, &data.Program)
	if err != nil {
		return nil, err
	}
	return *out, nil
}
