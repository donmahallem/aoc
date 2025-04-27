package day21_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day21"
)

const testData string = `029A
980A
179A
456A
379A`

func TestParseInput(t *testing.T) {
	raceWay := day21.ParseInput(strings.NewReader(testData))
	if len(*raceWay) != 5 {
		t.Errorf(`Expected %d to match 5`, len(*raceWay))
	}
}
