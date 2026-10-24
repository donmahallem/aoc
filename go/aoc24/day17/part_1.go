package day17

import (
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// maxSteps to prevent infinite runtime
const maxSteps = 1 << 16

// shiftRight returns value / 2^shift. Shifts of 63 or more return 0 instead of overflowing
func shiftRight(value, shift int) int {
	if shift <= 0 {
		return value
	}
	if shift >= 63 {
		return 0
	}
	return value / (1 << shift)
}

func CalculateOutput(reg *Register, program *Program) (*Program, error) {
	output := make(Program, 0)
	pointer := 0
	steps := 0
	var opcode, operand, operand_value int
	if len(*program)%2 != 0 {
		return nil, aoc_utils.NewParseError("malformed program: odd length", nil)
	}
	for pointer < len(*program) {
		steps++
		if steps > maxSteps {
			// circuit breaker
			return nil, aoc_utils.NewSolverErrorf("program did not halt within %d steps", maxSteps)
		}
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
			(*reg)[0] = shiftRight((*reg)[0], operand_value)
		case 1:
			(*reg)[1] = (*reg)[1] ^ operand
		case 2:
			(*reg)[1] = operand_value % 8
		case 3:
			if (*reg)[0] != 0 {
				if operand >= len(*program) {
					return nil, aoc_utils.NewParseError("invalid jump target", nil)
				}
				pointer = operand
				continue
			}
		case 4:
			(*reg)[1] = (*reg)[1] ^ (*reg)[2]
		case 5:
			output = append(output, operand_value%8)
		case 6:
			(*reg)[1] = shiftRight((*reg)[0], operand_value)
		case 7:
			(*reg)[2] = shiftRight((*reg)[0], operand_value)
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
