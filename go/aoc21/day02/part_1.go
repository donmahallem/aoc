package day02

import (
	"bufio"
	"io"
)

func Part1(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	posX, posY := 0, 0

	for s.Scan() {
		dir, val, err := parseLine(s.Bytes())
		if err != nil {
			return 0, err
		}
		switch dir {
		case Forward:
			posX += val
		case Down:
			posY += val
		case Up:
			posY -= val
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}
	return posX * posY, nil
}
