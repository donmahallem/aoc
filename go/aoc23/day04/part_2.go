package day04

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/aoc_utils"
)

func CountTickets(valid *[]int, idx int) int {
	counter := make(map[int]int)
	add := func(idx, count int) {
		//fmt.Printf("Add %d to %d. Was %d\n", count, idx, counter[idx])
		if val, ok := counter[idx]; ok {
			counter[idx] = val + count
		} else {
			counter[idx] = count
		}
	}
	for i := range *valid {
		for j := i + 1; j < aoc_utils.Min(i+1+(*valid)[i], len(*valid)); j++ {
			add(j, 1+counter[i])
		}
		add(i, 1)
	}
	//fmt.Printf("%v\n", counter)
	score := 0
	for key := range counter {
		score += counter[key]
	}
	return score
}
func Part2(in io.Reader) int {
	s := bufio.NewScanner(in)
	valid := make([]int, 0)
	for s.Scan() {
		_, a, b := ParseLine(s.Bytes())
		valid = append(valid, CountWinnings(a, b))
	}
	return CountTickets(&valid, 0)
}
