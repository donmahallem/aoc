package day22

import (
	"bufio"
	"io"
	"strconv"
)

const pruneValue int = 16777216                    // 2^24
const pruneValueMask int32 = int32(pruneValue - 1) // 2^24-1
func Step(secret int32) int32 {
	tmp := (secret ^ (secret << 6)) & pruneValueMask
	tmp = (tmp ^ (tmp >> 5)) & pruneValueMask
	tmp = (tmp ^ (tmp << 11)) & pruneValueMask
	return tmp
}

func SimulateSteps(secret int32, iterations uint16) int32 {
	tmp := secret
	for range iterations {
		tmp = Step(tmp)
	}
	return tmp
}

func ParseInput(in io.Reader) *[]int32 {
	points := make([]int32, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		tmp, _ := strconv.Atoi(line)
		points = append(points, int32(tmp))
	}
	return &points
}

const iterations uint16 = 2000

func AddUpSecrets(secrets *[]int32) int32 {
	sum := int32(0)
	for _, secret := range *secrets {
		sum += SimulateSteps(secret, iterations)
	}
	return sum
}

func Part1(in io.Reader) int32 {
	items := ParseInput(in)
	return AddUpSecrets(items)
}
