package day17

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type Register = [3]int
type Program = []int
type InputData struct {
	Register Register
	Program  Program
}

func parseInput(in io.Reader) (*InputData, error) {
	reg := Register{}
	var prog Program
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		if strings.TrimSpace(line) == "" {
			// skip blank lines
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		if parts[0] == "Register" {
			if len(parts) < 3 {
				return nil, aoc_utils.NewUnexpectedInputError(0)
			}
			regIdx := int(parts[1][0] - 'A')
			if regIdx < 0 || regIdx >= len(reg) {
				return nil, aoc_utils.NewUnexpectedInputError(parts[1][0])
			}
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, aoc_utils.NewParseError("invalid register value", err)
			}
			reg[regIdx] = val
		} else if strings.HasPrefix(parts[0], "Program:") {
			if len(parts) < 2 {
				return nil, aoc_utils.NewUnexpectedInputError(0)
			}
			splitOperands := strings.Split(parts[1], ",")
			prog = make(Program, len(splitOperands))
			for idx := range splitOperands {
				op := strings.TrimSpace(splitOperands[idx])
				if op == "" {
					return nil, aoc_utils.NewUnexpectedInputError(0)
				}
				val, err := strconv.Atoi(op)
				if err != nil {
					return nil, aoc_utils.NewParseError("invalid program operand", err)
				}
				prog[idx] = val
			}
		} else {
			return nil, aoc_utils.NewUnexpectedInputError(line[0])
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return &InputData{
		Register: reg,
		Program:  prog,
	}, nil
}
