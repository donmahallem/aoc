package day21

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample_1.txt
var testDataSample1 string

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testDataSample1)
		res, err := parseInput(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if res.Width != 11 || res.Height != 11 {
			t.Errorf(`Expected width and height to be 11, got %d and %d`, res.Width, res.Height)
		}
		if res.StartX != 5 || res.StartY != 5 {
			t.Errorf(`Expected start position to be (5,5), got (%d,%d)`, res.StartX, res.StartY)
		}

	})

}
