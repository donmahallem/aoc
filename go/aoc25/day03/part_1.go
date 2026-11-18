package day03

import (
	"bufio"
	"io"
)

const baseOffset = '0'

func Part1(in io.Reader) (int, error) {
	accumulator := 0
	scanner := bufio.NewScanner(in)

	var maxValueLeft byte
	var maxValueRight byte
	var foundIndex int
	for scanner.Scan() {
		maxValueLeft = 0
		maxValueRight = 0
		row := scanner.Bytes()

		for i := 0; i < len(row)-1; i++ {
			v := row[i]
			if v > maxValueLeft {
				maxValueLeft = v
				foundIndex = i
			}
		}
		for i := foundIndex + 1; i < len(row); i++ {
			if row[i] > maxValueRight {
				maxValueRight = row[i]
			}
		}
		accumulator += int((maxValueLeft-baseOffset)*10 + (maxValueRight - baseOffset))
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	return accumulator, nil
}
