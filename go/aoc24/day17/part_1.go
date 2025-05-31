package day17

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils/math/pow"
)

type Register = [3]int
type Program = []int

func ParseInput(in io.Reader) (Register, Program) {
	reg := Register{}
	var prog Program
	s := bufio.NewScanner(in)
	splitLine := make([]string, 0)
	for s.Scan() {
		line := s.Text()
		splitLine = append(splitLine[:0], strings.Split(line, " ")...)
		if splitLine[0] == "Register" {
			reg[byte(splitLine[1][0])-'A'], _ = strconv.Atoi(splitLine[2])
		} else if strings.HasPrefix(splitLine[0], "Program:") {
			splitOperands := strings.Split(splitLine[1], ",")
			prog = make(Program, len(splitOperands))
			for idx := range len(splitOperands) {
				prog[idx], _ = strconv.Atoi(splitOperands[idx])
			}
		}

	}
	return reg, prog
}

func CalculateOutput(reg *Register, program *Program) *Program {
	output := make(Program, 0)
	pointer := 0
	var opcode, operand, operand_value int
	for pointer < len(*program) {
		opcode = (*program)[pointer]
		operand = (*program)[pointer+1]
		if operand < 4 {
			operand_value = operand
		} else if operand >= 4 && operand < 7 {
			operand_value = (*reg)[operand-4]
		}
		switch opcode {
		case 0:
			(*reg)[0] = (*reg)[0] / pow.IntPow(2, operand_value)
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
			(*reg)[1] = (*reg)[0] / (pow.IntPow(2, operand_value))
		case 7:
			(*reg)[2] = (*reg)[0] / (pow.IntPow(2, operand_value))
		}
		pointer += 2
	}
	return &output
}

func Part1(in io.Reader) []int {
	registers, program := ParseInput(in)
	return *CalculateOutput(&registers, &program)
}
