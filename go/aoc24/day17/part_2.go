package day17

import (
	"io"
	"math"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

// maxSearchRuns bounds the number of candidate register values tried by A.
const maxSearchRuns = 1 << 20

func SearchForNeutralElement(reg *Register, program *Program, targetProgram *Program) (bool, error) {
	var outputBuf [16]int
	output := outputBuf[:0]
	pointer := 0
	steps := 0
	targetProgramLength := len(*targetProgram)
	var opcode, operand, operand_value, programLength int
	if len(*program)%2 != 0 {
		return false, aoc_utils.NewParseError("malformed program: odd length", nil)
	}
	for pointer < len(*program) {
		steps++
		if steps > maxSteps {
			return false, aoc_utils.NewSolverErrorf("program did not halt within %d steps", maxSteps)
		}
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
			(*reg)[0] = shiftRight((*reg)[0], operand_value)
		case 1:
			(*reg)[1] = (*reg)[1] ^ operand
		case 2:
			(*reg)[1] = operand_value % 8
		case 3:
			if (*reg)[0] != 0 {
				if operand >= len(*program) {
					return false, aoc_utils.NewParseError("invalid jump target", nil)
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
			return false, aoc_utils.NewParseError("invalid opcode", nil)
		}
		pointer += 2
	}
	return slices.Equal(*targetProgram, output), nil
}

type searcher struct {
	program *Program
	runs    int
}

func (s *searcher) find(prefix, level int) (int, bool, error) {
	target := (*s.program)[level:]
	for digit := 0; digit < 8; digit++ {
		if prefix > (math.MaxInt-digit)/8 {
			continue
		}
		candidate := prefix*8 + digit
		s.runs++
		if s.runs > maxSearchRuns {
			return 0, false, aoc_utils.NewSolverErrorf("search budget of %d runs exhausted", maxSearchRuns)
		}
		reg := Register{candidate, 0, 0}
		found, err := SearchForNeutralElement(&reg, s.program, &target)
		if err != nil {
			return 0, false, err
		}
		if !found {
			continue
		}
		if level == 0 {
			return candidate, true, nil
		}
		value, ok, err := s.find(candidate, level-1)
		if err != nil || ok {
			return value, ok, err
		}
	}
	return 0, false, nil
}

func A(program *Program) (int, error) {
	if len(*program) == 0 {
		return 0, nil
	}
	s := searcher{program: program}
	value, ok, err := s.find(0, len(*program)-1)
	if err != nil {
		return 0, err
	}
	if !ok {
		return 0, aoc_utils.NewSolverErrorf("no register A reproduces the program")
	}
	return value, nil
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
