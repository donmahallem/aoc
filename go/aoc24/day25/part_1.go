package day25

import (
	"bufio"
	"io"
)

type InputType = byte

const (
	InputTypeLock InputType = '#'
	InputTypeKey  InputType = '.'
	InputTypeNone InputType = 0
)

func Part1(in io.Reader) (uint32, error) {
	scanner := bufio.NewScanner(in)

	locks := make([]uint32, 0, 256)
	keys := make([]uint32, 0, 256)

	currentType := InputTypeNone
	currentEncoded := uint32(0)
	currentLine := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) < 5 {
			// blank line. Block ends
			switch currentType {
			case InputTypeLock:
				locks = append(locks, currentEncoded)
			case InputTypeKey:
				keys = append(keys, currentEncoded)
			}
			currentType = InputTypeNone
			currentEncoded = 0
			currentLine = 0
			continue
		}

		switch currentLine {
		case 0:
			currentType = line[0]
		case 6:
			// footer row, carries no information
		default:
			row := line[:5]
			offset := currentLine - 1
			for i := range 5 {
				if row[i] == InputTypeLock {
					currentEncoded |= 1 << (offset + (4-i)*5)
				}
			}
		}
		currentLine++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	switch currentType {
	case InputTypeLock:
		locks = append(locks, currentEncoded)
	case InputTypeKey:
		keys = append(keys, currentEncoded)
	}

	count := uint32(0)
	for _, l := range locks {
		for _, k := range keys {
			if l&k == 0 {
				count++
			}
		}
	}
	return count, nil
}
