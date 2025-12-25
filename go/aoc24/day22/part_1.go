package day22

import (
	"bufio"
	"io"
	"strconv"
)

const pruneValue uint32 = 16777216           // 2^24
const pruneValueMask uint32 = pruneValue - 1 // 2^24-1
func Step(secret uint32) uint32 {
	tmp := (secret ^ (secret << 6)) & pruneValueMask
	tmp = (tmp ^ (tmp >> 5)) & pruneValueMask
	tmp = (tmp ^ (tmp << 11)) & pruneValueMask
	return tmp
}

func SimulateSteps(secret uint32, iterations uint16) uint32 {
	tmp := secret
	for range iterations {
		tmp = Step(tmp)
	}
	return tmp
}

func ParseInput(in io.Reader) *[]uint32 {
	points := make([]uint32, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		tmp, _ := strconv.Atoi(line)
		points = append(points, uint32(tmp))
	}
	return &points
}

const iterations uint16 = 2000

func AddUpSecrets(secrets *[]uint32) uint {
	sum := uint(0)
	for _, secret := range *secrets {
		sum += uint(SimulateSteps(secret, iterations))
	}
	return sum
}

func Part1(in io.Reader) (uint, error) {
	items := ParseInput(in)
	return AddUpSecrets(items), nil
}
