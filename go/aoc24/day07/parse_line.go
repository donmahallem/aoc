package day07

import "github.com/donmahallem/aoc/go/aoc_utils"

type parsedLineData struct {
	Result     int
	TestValues []int
}

func parseLine(raw []byte, parsed *parsedLineData) error {
	parsed.Result = 0
	parsed.TestValues = parsed.TestValues[:0]

	// State: 0 = Parsing Target, 1 = Parsing Values
	state := 0
	currentVal := 0
	hasVal := false

	for _, c := range raw {
		if state == 0 {
			if c >= '0' && c <= '9' {
				parsed.Result = parsed.Result*10 + int(c-'0')
			} else if c == ':' {
				state = 1
			} else {
				return aoc_utils.NewParseError("expected colon after target", nil)
			}
		} else {
			if c == ' ' {
				if hasVal {
					parsed.TestValues = append(parsed.TestValues, currentVal)
					currentVal = 0
					hasVal = false
				}
			} else if c >= '0' && c <= '9' {
				currentVal = currentVal*10 + int(c-'0')
				hasVal = true
			} else {
				return aoc_utils.NewParseError("invalid character in values", nil)
			}
		}
	}

	if state == 0 {
		return aoc_utils.NewParseError("missing colon", nil)
	}

	if hasVal {
		parsed.TestValues = append(parsed.TestValues, currentVal)
	}
	return nil
}
