package day22_test

import (
	"fmt"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day22"
)

var testCases []int = []int{
	123,
	15887950,
	16495136,
	527345,
	704524,
	1553684,
	12683156,
	11100544,
	12249484,
	7753432,
	5908254}

func TestStep(t *testing.T) {
	for idx := range len(testCases) - 1 {
		t.Run(fmt.Sprintf("Expect %d to become %d", testCases[idx], testCases[idx+1]), func(t *testing.T) {
			points := day22.Step(uint32(testCases[idx]))
			if points != uint32(testCases[idx+1]) {
				t.Errorf(`Expected %d to match %d`, points, testCases[idx+1])
			}
		})
	}
}
