package day05

import (
	"bufio"
	"fmt"
	"io"
	"sort"

	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

// tryFixLine reorders the update to satisfy all rules and returns its median entry.
func tryFixLine(f facts, line []int64) (int64, bool) {
	sort.SliceStable(line, func(i, j int) bool {
		a := line[i]
		b := line[j]
		pairBA := encodePair(b, a)
		if _, ok := f[encodePair(a, b)]; ok {
			if _, reverse := f[pairBA]; reverse {
				return a < b
			}
			return true
		}
		if _, ok := f[pairBA]; ok {
			return false
		}
		return a < b
	})
	return validateLine(f, line)
}

func Part2(in io.Reader) int64 {
	s := bufio.NewScanner(in)
	baseData := true
	rules := make(facts)
	var numberA, numberB int64
	total := int64(0)
	numbers := make([]int64, 0, 16)

	for s.Scan() {
		lineData := s.Bytes()
		if len(lineData) == 0 {
			baseData = false
			continue
		}
		if baseData {
			numberA, numberB = 0, 0
			currentNumber := &numberA
			for _, c := range lineData {
				if bytes.ByteIsNumber(c) {
					*currentNumber = (*currentNumber)*10 + int64(c-'0')
				} else if c == '|' {
					currentNumber = &numberB
				} else {
					panic(fmt.Sprintf("Unexpected character %c in base data", c))
				}
			}
			rules[encodePair(numberA, numberB)] = struct{}{}
			continue
		}

		bytes.ParseIntSequence(lineData, ',', &numbers)
		if len(numbers) == 0 {
			continue
		}
		if _, ok := validateLine(rules, numbers); ok {
			continue
		}
		if mid, ok := tryFixLine(rules, numbers); ok {
			total += mid
		}
	}
	return total
}
