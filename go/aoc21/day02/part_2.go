package day02

import (
	"bufio"
	"io"
)

func Part2(in io.Reader) (int, error) {
	s := bufio.NewScanner(in)
	posX, posY, aim := 0, 0, 0
	for s.Scan() {
		dir, val, err := parseLine(s.Bytes())
		if err != nil {
			return 0, err
		}
		switch dir {
		case Forward:
			posX += val
			posY += aim * val
		case Down:
			aim += val
		case Up:
			aim -= val
		}
	}
	if s.Err() != nil {
		return 0, s.Err()
	}
	return posX * posY, nil
}
