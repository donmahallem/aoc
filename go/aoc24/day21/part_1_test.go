package day21_test

import (
	"fmt"
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
	fmt.Printf("%v", raceWay)
}

func TestWalk_Depth1(t *testing.T) {
	t.Run("A to 0", func(t *testing.T) {
		start := day21.NUMERIC_A
		end := day21.NUMERIC_0
		_, result := day21.Walk(&start, &end, 0, 1)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("0 to 2", func(t *testing.T) {
		start := day21.NUMERIC_0
		end := day21.NUMERIC_2
		_, result := day21.Walk(&start, &end, 0, 1)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("2 to 9", func(t *testing.T) {
		start := day21.NUMERIC_2
		end := day21.NUMERIC_9
		_, result := day21.Walk(&start, &end, 0, 1)
		if result != 4 {
			t.Errorf(`Expected %d to match 4`, result)
		}
	})
	t.Run("9 to A", func(t *testing.T) {
		start := day21.NUMERIC_9
		end := day21.NUMERIC_A
		_, result := day21.Walk(&start, &end, 0, 1)
		if result != 4 {
			t.Errorf(`Expected %d to match 4`, result)
		}
	})
	t.Run("A to 7", func(t *testing.T) {
		start := day21.NUMERIC_A
		end := day21.NUMERIC_7
		_, result := day21.Walk(&start, &end, 0, 1)
		if result != 6 {
			t.Errorf(`Expected %d to match 5`, result)
		}
	})
}

func TestWalk_Depth2(t *testing.T) {
	t.Run("A to 0", func(t *testing.T) {
		start := day21.NUMERIC_A
		end := day21.NUMERIC_0
		_, result := day21.Walk(&start, &end, 0, 2)
		if result != 8 {
			t.Errorf(`Expected %d to match 8`, result)
		}
	})
	t.Run("0 to 2", func(t *testing.T) {
		start := day21.NUMERIC_0
		end := day21.NUMERIC_2
		_, result := day21.Walk(&start, &end, 0, 2)
		if result != 4 {
			t.Errorf(`Expected %d to match 4`, result)
		}
	})
	t.Run("2 to 9", func(t *testing.T) {
		start := day21.NUMERIC_2
		end := day21.NUMERIC_9
		_, result := day21.Walk(&start, &end, 0, 2)
		if result != 8 {
			t.Errorf(`Expected %d to match 8`, result)
		}
	})
	t.Run("9 to A", func(t *testing.T) {
		start := day21.NUMERIC_9
		end := day21.NUMERIC_A
		_, result := day21.Walk(&start, &end, 0, 2)
		if result != 8 {
			t.Errorf(`Expected %d to match 8`, result)
		}
	})
}
