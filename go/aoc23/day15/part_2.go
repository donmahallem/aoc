package day15

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

func Part2(in io.Reader) (uint64, error) {
	// Stream-parse input to avoid one large allocation — but pre-allocate small per-bucket capacity to reduce realloc churn.
	scanner := bufio.NewScanner(in)
	key := make([]byte, 0, 16)
	readingValue := false
	valueRead := false
	valAcc := byte(0)
	removal := false

	memory := make([][]lens, 256)

	var handleToken = func() error {
		if len(key) == 0 {
			return aoc_utils.NewParseError("empty key", nil)
		}
		idHash, id := hashId(key)
		if removal {
			box := memory[idHash]
			for idx, l := range box {
				if l.Id == id {
					n := len(box)
					if idx < n-1 {
						copy(box[idx:], box[idx+1:])
					}
					box[n-1] = lens{}
					memory[idHash] = box[:n-1]
					return nil
				}
			}
			return nil
		}
		if !valueRead {
			return aoc_utils.NewParseError("missing value", nil)
		}
		// update or append; pre-allocate small capacity on first use to avoid repeated reallocs
		if memory[idHash] == nil {
			memory[idHash] = make([]lens, 0, 8)
		}
		box := memory[idHash]
		for idx, l := range box {
			if l.Id == id {
				memory[idHash][idx].FocalLength = uint64(valAcc)
				return nil
			}
		}
		memory[idHash] = append(memory[idHash], lens{Id: id, FocalLength: uint64(valAcc)})
		return nil
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		readingValue = false
		valueRead = false
		valAcc = 0
		removal = false
		key = key[:0]

		for _, ch := range line {
			switch ch {
			case ',':
				if err := handleToken(); err != nil {
					return 0, err
				}
				readingValue = false
				valueRead = false
				valAcc = 0
				removal = false
				key = key[:0]
			case '=':
				readingValue = true
			case '-':
				if readingValue {
					return 0, aoc_utils.NewParseError("unexpected '-' in value", nil)
				}
				removal = true
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !readingValue {
					return 0, aoc_utils.NewParseError("digit in key", nil)
				}
				if !valueRead {
					valueRead = true
					valAcc = ch - '0'
				} else {
					return 0, aoc_utils.NewParseError("value too long", nil)
				}
			default:
				if readingValue {
					return 0, aoc_utils.NewParseError("invalid character in value", nil)
				}
				key = append(key, ch)
			}
		}
		if len(key) > 0 || valueRead || removal {
			if err := handleToken(); err != nil {
				return 0, err
			}
		}
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	var totalFocusingPower uint64 = 0
	for boxIdx, box := range memory {
		for lensIdx, lense := range box {
			totalFocusingPower += (uint64(boxIdx) + 1) * (uint64(lensIdx) + 1) * lense.FocalLength
		}
	}
	return totalFocusingPower, nil
}
