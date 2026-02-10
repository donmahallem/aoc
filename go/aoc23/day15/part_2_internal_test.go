package day15

import (
	"strings"
	"testing"
)

var testData string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestParseInput(t *testing.T) {
	t.Run("test parse input", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := parseInput(reader)
		expectedLen := 11
		if len(res) != expectedLen {
			t.Errorf(`Expected number of groups to be %d, got %d`, expectedLen, len(res))
		}
	})
}
