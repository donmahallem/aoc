package day15_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day15"
)

const testData string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
const testDataHash string = `HASH`

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {

		reader := strings.NewReader(testDataHash)
		res := day15.Part1(reader)
		if res != 52 {
			t.Errorf(`Expected number of blocks to be 405, got %d`, res)
		}
	})
	t.Run("test sample 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day15.Part1(reader)
		if res != 1320 {
			t.Errorf(`Expected number of blocks to be 405, got %d`, res)
		}
	})
}
