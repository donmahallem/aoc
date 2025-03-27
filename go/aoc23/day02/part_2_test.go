package day02_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc23/day02"
)

func TestCalculateMinBlock(t *testing.T) {
	data := []byte(testData[0])
	_, blocks := day02.ParseLine(&data)
	if res := day02.CalculateMinBlock(&blocks); res != 48 {
		t.Errorf(`Expected %d to be %d`, res, 48)
	}
}
