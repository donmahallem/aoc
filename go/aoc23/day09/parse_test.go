package day09

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res, _ := parseInput(reader)
		if len(res) != 3 {
			t.Errorf(`Expected 3 rows not %d`, len(res))
		}
		if len(res[0]) != 6 {
			t.Errorf(`Expected 6 numbers not %d`, res[0])
		}
	})
}
