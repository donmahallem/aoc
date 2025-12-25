package day19

import (
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		data := parseInput(strings.NewReader(sample))
		if len(data.Workflows) != 11 {
			t.Errorf("len(data.Workflows) = %v, want %v", len(data.Workflows), 11)
		}
		if len(data.Ratings) != 5 {
			t.Errorf("len(data.Ratings) = %v, want %v", len(data.Ratings), 5)
		}
	})
}
