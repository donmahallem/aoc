package day04

import (
	"bufio"
	"io"
)

func CountTickets(valid []int, idx int) int {
	counter := make(map[int]int)
	add := func(idx, count int) {
		if val, ok := counter[idx]; ok {
			counter[idx] = val + count
		} else {
			counter[idx] = count
		}
	}
	for i := range len(valid) {
		for j := i + 1; j < min(i+1+valid[i], len(valid)); j++ {
			add(j, 1+counter[i])
		}
		add(i, 1)
	}
	score := 0
	for key := range counter {
		score += counter[key]
	}
	return score
}
func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	valid := make([]int, 0, 32)
	for s.Scan() {
		_, a, b := ParseLine(s.Bytes())
		valid = append(valid, CountWinnings(a, b))
	}
	return CountTickets(valid, 0)
}
